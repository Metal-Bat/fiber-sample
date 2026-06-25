package user

import (
	"context"
	"errors"
	"time"

	"sample/src/initializers"
	"sample/src/pkg/base"
	"sample/src/pkg/dto"
	"sample/src/pkg/mapper"
	"sample/src/pkg/utils"
	"sample/src/pkg/utils/customErrors"
)

type Service interface {
	base.Service[dto.UserInfo, dto.UserDetail, dto.CreateUser, dto.UpdateUser]
	Login(ctx context.Context, mobile string, password string) (*dto.LoginResult, error)
}

type userService struct {
	base.Service[dto.UserInfo, dto.UserDetail, dto.CreateUser, dto.UpdateUser]
	userRepo UserRepository
	mapper   mapper.UserConverter
}

func NewUserService(
	baseSvc base.Service[dto.UserInfo, dto.UserDetail, dto.CreateUser, dto.UpdateUser],
	userRepo UserRepository,
	m mapper.UserConverter,
) Service {
	return &userService{
		Service:  baseSvc,
		userRepo: userRepo,
		mapper:   m,
	}
}

func (s *userService) Login(ctx context.Context, mobile string, password string) (*dto.LoginResult, error) {
	_, span := initializers.Tracer.Start(ctx, "service.User.Login")
	defer span.End()

	user, err := s.userRepo.FindByMobileWithPermissions(ctx, mobile)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	if err := utils.ComparHash(ctx, user.Password, password); err != nil {
		span.RecordError(err)
		return nil, &customErrors.InternalError{
			HTTPStatus: 401,
			MessageKey: "invalid user or password",
			Err:        errors.New("invalid credentials"),
		}
	}

	expiresAt := time.Now().Add(time.Hour * 72).Unix()
	permissions := make([]string, len(user.Permissions))
	for i, p := range user.Permissions {
		permissions[i] = p.Name
	}

	token, err := utils.CreateJwtToken(ctx, mobile, permissions, expiresAt)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	return &dto.LoginResult{
		UserInfo:  *s.mapper.ToList(user),
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}
