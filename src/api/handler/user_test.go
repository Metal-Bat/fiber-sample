package handler

import (
	"context"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"sample/src/pkg/base"
	"sample/src/pkg/dto"
)

type mockUserService struct {
	mock.Mock
}

func (m *mockUserService) Name() string { return "User" }

func (m *mockUserService) GetAll(ctx context.Context, page dto.PaginationStructure) ([]*dto.UserInfo, int64, error) {
	args := m.Called()
	return args.Get(0).([]*dto.UserInfo), args.Get(1).(int64), args.Error(2)
}

func (m *mockUserService) GetOne(ctx context.Context, index uint) (*dto.UserDetail, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.UserDetail), args.Error(1)
}

func (m *mockUserService) Create(ctx context.Context, data *dto.CreateUser) (*dto.UserDetail, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.UserDetail), args.Error(1)
}

func (m *mockUserService) Update(ctx context.Context, index uint, data *dto.UpdateUser) error {
	args := m.Called()
	return args.Error(0)
}

func (m *mockUserService) Delete(ctx context.Context, index uint) error {
	args := m.Called()
	return args.Error(0)
}

func (m *mockUserService) Login(ctx context.Context, mobile string, password string) (*dto.LoginResult, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.LoginResult), args.Error(1)
}

func TestGetUsers_Success(t *testing.T) {
	app := fiber.New()
	svc := new(mockUserService)

	svc.On("GetAll", mock.Anything, mock.Anything).Return(
		[]*dto.UserInfo{{ID: 1, Mobile: "0912", Email: "test@test.com", Permissions: []string{}}},
		int64(1),
		nil,
	)

	var crud base.Service[dto.UserInfo, dto.UserDetail, dto.CreateUser, dto.UpdateUser] = svc
	app.Get("/users", GetAll(crud, []string{"id", "mobile"}))

	req := httptest.NewRequest("GET", "/users?page=1&size=10", nil)
	resp, _ := app.Test(req, fiber.TestConfig{Timeout: 0})
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Contains(t, string(body), "result")
	assert.Contains(t, string(body), "success")
}

func TestGetUser_Success(t *testing.T) {
	app := fiber.New()
	svc := new(mockUserService)

	svc.On("GetOne", mock.Anything, mock.Anything).Return(
		&dto.UserDetail{ID: 1, Mobile: "0912", Email: "test@test.com"},
		nil,
	)

	var crud base.Service[dto.UserInfo, dto.UserDetail, dto.CreateUser, dto.UpdateUser] = svc
	app.Get("/users/:id", GetOne(crud))

	req := httptest.NewRequest("GET", "/users/1", nil)
	resp, _ := app.Test(req, fiber.TestConfig{Timeout: 0})
	assert.Equal(t, 200, resp.StatusCode)
}
