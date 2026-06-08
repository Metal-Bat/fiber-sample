package user

import (
	"errors"
	"sample/src/initializers"
	"sample/src/pkg/dto"
	"sample/src/pkg/entities"
	"sample/src/pkg/utils"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type UserRepository interface {
	// crud
	GetUsersCount(c fiber.Ctx, page dto.PaginationStructure) (int64, error)
	GetUsers(c fiber.Ctx, page dto.PaginationStructure) ([]*entities.User, error)
	GetUser(c fiber.Ctx, index uint) (*entities.User, error)
	CreateUser(c fiber.Ctx, user *entities.User) (*entities.User, error)
	UpdateUser(c fiber.Ctx, user *entities.User) error
	DeleteUser(c fiber.Ctx, user *entities.User) error
	// service
	FindByMobileWithPermissions(c fiber.Ctx, mobile string) (*entities.User, error)
}

type repository struct {
	DB *gorm.DB
}

func NewUserRepository() UserRepository {
	return &repository{
		initializers.DB,
	}
}

func (r *repository) GetUsersCount(
	c fiber.Ctx,
	page dto.PaginationStructure,
) (int64, error) {
	_, span := initializers.Tracer.Start(c.Context(), "repository.GetUsers")
	defer span.End()

	q := gorm.G[entities.User](r.DB.WithContext(c.Context())).Order("id")
	count, err := utils.ApplyFilters(c, q, page).Count(c.Context(), "id")

	if err != nil {
		span.RecordError(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("user count error")
		}
		return 0, err
	}

	return count, nil
}

func (r *repository) GetUsers(
	c fiber.Ctx,
	page dto.PaginationStructure,
) ([]*entities.User, error) {

	_, span := initializers.Tracer.Start(c.Context(), "repository.GetUsers")
	defer span.End()

	q := gorm.G[entities.User](r.DB.WithContext(c.Context())).Order("id")
	q = utils.ApplyFilters(c, q, page)
	users, err := utils.ApplyPagination(c, q, page).Find(c.Context())

	if err != nil {
		span.RecordError(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("users not found")
		}
		return nil, err
	}

	result := make([]*entities.User, len(users))
	for i := range users {
		result[i] = &users[i]
	}

	return result, nil
}

func (r *repository) GetUser(c fiber.Ctx, index uint) (*entities.User, error) {
	_, span := initializers.Tracer.Start(c.Context(), "repository.GetUser")
	defer span.End()

	user, err := gorm.G[entities.User](r.DB.WithContext(c.Context())).
		Preload(
			"Permissions",
			func(db gorm.PreloadBuilder) error {
				return nil
			},
		).
		Where("id = ?", index).
		First(c.Context())

	if err != nil {
		span.RecordError(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *repository) CreateUser(c fiber.Ctx, user *entities.User) (*entities.User, error) {
	_, span := initializers.Tracer.Start(c.Context(), "repository.CreateUser")
	defer span.End()

	err := gorm.G[entities.User](r.DB.WithContext(c.Context())).Create(c.Context(), user)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	return user, nil
}

func (r *repository) UpdateUser(c fiber.Ctx, user *entities.User) error {
	_, span := initializers.Tracer.Start(c.Context(), "repository.UpdateUser")
	defer span.End()

	_, err := gorm.G[entities.User](r.DB.WithContext(c.Context())).
		Where("id = ?", user.ID).
		Updates(c.Context(), *user)

	if err != nil {
		span.RecordError(err)
		return err
	}
	return nil
}

func (r *repository) DeleteUser(c fiber.Ctx, user *entities.User) error {
	_, span := initializers.Tracer.Start(c.Context(), "repository.DeleteUser")
	defer span.End()

	_, err := gorm.G[entities.User](r.DB.WithContext(c.Context())).
		Where("id = ?", user.ID).
		Delete(c.Context())

	if err != nil {
		span.RecordError(err)
		return err
	}
	return nil
}

func (r *repository) FindByMobileWithPermissions(
	c fiber.Ctx,
	mobile string,
) (*entities.User, error) {
	_, span := initializers.Tracer.Start(c.Context(), "repository.FindByMobileWithPermissions")
	defer span.End()

	user, err := gorm.G[entities.User](r.DB.WithContext(c.Context())).
		Preload(
			"Permissions",
			func(db gorm.PreloadBuilder) error {
				return nil
			},
		).
		Where("mobile = ?", mobile).
		First(c.Context())

	if err != nil {
		span.RecordError(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
