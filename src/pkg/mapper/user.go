package mapper

import (
	"database/sql"
	"sample/src/initializers"
	"sample/src/pkg/dto"
	"sample/src/pkg/entities"
	"sample/src/pkg/utils"

	"github.com/gofiber/fiber/v3"
)

type UserMapper interface {
	// crud
	ToUserDto(c fiber.Ctx, u *entities.User) *dto.UserInfo
	ToUserListDto(c fiber.Ctx, users []*entities.User) []*dto.UserInfo
	ToUserDetailDto(c fiber.Ctx, user *entities.User) *dto.UserDetail
	ToCreateDto(c fiber.Ctx, data *dto.CreateUser) *entities.User
	ToUpdateDto(c fiber.Ctx, data *dto.UpdateUser, user *entities.User)
	// service
	ToLoginDto(c fiber.Ctx, user entities.User, token string, expiresAt int64) *dto.LoginResult
}

type userMapper struct{}

func NewUserMapper() UserMapper {
	return &userMapper{}
}

func (m *userMapper) ToUserDto(
	c fiber.Ctx,
	user *entities.User,
) *dto.UserInfo {
	_, span := initializers.Tracer.Start(c.Context(), "mapper.UserEntityToUserDto")
	defer span.End()

	permissions := make([]string, len(user.Permissions))
	for i, p := range user.Permissions {
		permissions[i] = p.Name
	}

	return &dto.UserInfo{
		ID:           user.ID,
		Mobile:       user.Mobile,
		Email:        user.Email,
		NationalCode: utils.NullableString(user.NationalCode),
		BirthDate:    utils.NullableTime(user.BirthDate),
		Permissions:  permissions,
	}
}

func (m *userMapper) ToUserDetailDto(
	c fiber.Ctx,
	user *entities.User,
) *dto.UserDetail {
	_, span := initializers.Tracer.Start(c.Context(), "mapper.UserEntityToUserDetailDto")
	defer span.End()

	permissions := make([]string, len(user.Permissions))
	for i, p := range user.Permissions {
		permissions[i] = p.Name
	}

	return &dto.UserDetail{
		ID:           user.ID,
		Mobile:       user.Mobile,
		Email:        user.Email,
		NationalCode: utils.NullableString(user.NationalCode),
		BirthDate:    utils.NullableTime(user.BirthDate),
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		DeletedAt:    utils.NullableTime(sql.NullTime(user.DeletedAt)),
		Permissions:  permissions,
	}
}

func (m *userMapper) ToUserListDto(
	c fiber.Ctx,
	users []*entities.User,
) []*dto.UserInfo {
	_, span := initializers.Tracer.Start(c.Context(), "service.Login")
	defer span.End()
	return MapSlice(c, users, m.ToUserDto)
}

func (m *userMapper) ToCreateDto(c fiber.Ctx, data *dto.CreateUser) *entities.User {

	user := &entities.User{
		Mobile:   data.Mobile,
		Email:    data.Email,
		Password: data.Password,
	}

	if data.NationalCode != nil {
		user.NationalCode = sql.NullString{
			String: *data.NationalCode,
			Valid:  true,
		}
	}

	if data.BirthDate != nil {
		user.BirthDate = sql.NullTime{
			Time:  *data.BirthDate,
			Valid: true,
		}
	}

	return user
}

func (m *userMapper) ToUpdateDto(c fiber.Ctx, data *dto.UpdateUser, user *entities.User) {

	if data.Email != nil {
		user.Email = *data.Email
	}

	if data.Password != nil {
		user.Password = *data.Password
	}

	if data.NationalCode != nil {
		user.NationalCode = sql.NullString{
			String: *data.NationalCode,
			Valid:  true,
		}
	}

	if data.BirthDate != nil {
		user.BirthDate = sql.NullTime{
			Time:  *data.BirthDate,
			Valid: true,
		}
	}
}

func (m *userMapper) ToLoginDto(
	c fiber.Ctx,
	user entities.User,
	token string,
	expiresAt int64,
) *dto.LoginResult {
	_, span := initializers.Tracer.Start(c.Context(), "mapper.LoginMapper")
	defer span.End()

	return &dto.LoginResult{
		UserInfo:  *m.ToUserDto(c, &user),
		Token:     token,
		ExpiresAt: expiresAt,
	}
}
