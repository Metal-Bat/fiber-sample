package user

import (
	"errors"
	"sample/src/initializers"
	"sample/src/pkg/dto"
	"sample/src/pkg/mapper"
	"sample/src/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v3"
)

type Service interface {
	// crud
	GetUsers(c fiber.Ctx, page dto.PaginationStructure) ([]*dto.UserInfo, int64, error)
	GetUser(c fiber.Ctx, index uint) (*dto.UserDetail, error)
	CreateUser(c fiber.Ctx, data *dto.CreateUser) (*dto.UserDetail, error)
	UpdateUser(c fiber.Ctx, index uint, data *dto.UpdateUser) error
	DeleteUser(c fiber.Ctx, index uint) error

	// service
	Login(c fiber.Ctx, mobile string, password string) (*dto.LoginResult, error)
}

type service struct {
	repository UserRepository
	mapper     mapper.UserMapper
}

func NewUserService(
	r UserRepository,
	m mapper.UserMapper,
) Service {
	return &service{
		repository: r,
		mapper:     m,
	}
}

func (s *service) GetUsers(
	c fiber.Ctx,
	page dto.PaginationStructure,
) ([]*dto.UserInfo, int64, error) {
	_, span := initializers.Tracer.Start(c.Context(), "service.GetUsers")
	defer span.End()
	users, err := s.repository.GetUsers(c, page)
	if err != nil {
		span.RecordError(err)
		return nil, 0, err
	}
	count, err := s.repository.GetUsersCount(c, page)
	if err != nil {
		span.RecordError(err)
		return nil, 0, err
	}

	return s.mapper.ToUserListDto(c, users), count, nil
}

func (s *service) GetUser(
	c fiber.Ctx,
	index uint,
) (*dto.UserDetail, error) {
	_, span := initializers.Tracer.Start(c.Context(), "service.GetUser")
	defer span.End()

	user, err := s.repository.GetUser(c, index)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	return s.mapper.ToUserDetailDto(c, user), nil
}

func (s *service) CreateUser(
	c fiber.Ctx,
	data *dto.CreateUser,
) (*dto.UserDetail, error) {
	_, span := initializers.Tracer.Start(c.Context(), "service.CreateUser")
	defer span.End()

	entity := s.mapper.ToCreateDto(c, data)

	hashed, err := utils.HashPassword(c, entity.Password)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	entity.Password = hashed

	user, err := s.repository.CreateUser(c, entity)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	return s.mapper.ToUserDetailDto(c, user), nil
}

func (s *service) UpdateUser(
	c fiber.Ctx,
	index uint,
	data *dto.UpdateUser,
) error {
	_, span := initializers.Tracer.Start(c.Context(), "service.UpdateUser")
	defer span.End()

	user, err := s.repository.GetUser(c, index)
	if err != nil {
		span.RecordError(err)
		return err
	}

	s.mapper.ToUpdateDto(c, data, user)

	if data.Password != nil && *data.Password != "" {
		hashed, err := utils.HashPassword(c, *data.Password)
		if err != nil {
			span.RecordError(err)
			return err
		}
		user.Password = hashed
	}

	err = s.repository.UpdateUser(c, user)
	if err != nil {
		span.RecordError(err)
		return err
	}

	return nil
}

func (s *service) DeleteUser(
	c fiber.Ctx,
	index uint,
) error {
	_, span := initializers.Tracer.Start(c.Context(), "service.DeleteUser")
	defer span.End()

	user, err := s.repository.GetUser(c, index)
	if err != nil {
		span.RecordError(err)
		return err
	}

	err = s.repository.DeleteUser(c, user)
	if err != nil {
		span.RecordError(err)
		return err
	}

	return nil
}

func (s *service) Login(c fiber.Ctx, mobile string, password string) (*dto.LoginResult, error) {
	_, span := initializers.Tracer.Start(c.Context(), "service.Login")
	defer span.End()

	user, err := s.repository.FindByMobileWithPermissions(c, mobile)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	if err := utils.ComparHash(c, user.Password, password); err != nil {
		span.RecordError(err)
		return nil, errors.New("invalid credentials")
	}

	duration := time.Now().Add(time.Hour * 72).Unix()
	permissions := make([]string, len(user.Permissions))
	for i, p := range user.Permissions {
		permissions[i] = p.Name
	}

	token, err := utils.CreateJwtToken(c, mobile, permissions, duration)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	return s.mapper.ToLoginDto(c, *user, token, duration), nil
}
