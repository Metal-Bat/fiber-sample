package handler

import (
	"io"
	"net/http/httptest"
	"sample/src/pkg/dto"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserService struct {
	mock.Mock
}

func (m *mockUserService) GetUsers(
	c fiber.Ctx,
	page dto.PaginationStructure,
) ([]*dto.UserInfo, int64, error) {
	args := m.Called()
	return args.Get(0).([]*dto.UserInfo), args.Get(1).(int64), args.Error(2)
}

func (m *mockUserService) GetUser(c fiber.Ctx, index uint) (*dto.UserDetail, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.UserDetail), args.Error(1)
}

func (m *mockUserService) CreateUser(c fiber.Ctx, data *dto.CreateUser) (*dto.UserDetail, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.UserDetail), args.Error(1)
}

func (m *mockUserService) UpdateUser(c fiber.Ctx, index uint, data *dto.UpdateUser) error {
	args := m.Called()
	return args.Error(0)
}

func (m *mockUserService) DeleteUser(c fiber.Ctx, index uint) error {
	args := m.Called()
	return args.Error(0)
}

func (m *mockUserService) Login(c fiber.Ctx, mobile string, password string) (*dto.LoginResult, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.LoginResult), args.Error(1)
}

func TestGetUsers_Success(t *testing.T) {
	app := fiber.New()
	svc := new(mockUserService)

	svc.On(
		"GetUsers",
		mock.Anything,
		mock.Anything,
	).
		Return(
			[]*dto.UserInfo{
				{
					ID:          1,
					Mobile:      "0912",
					Email:       "test@test.com",
					Permissions: []string{},
				},
			},
			int64(1),
			nil,
		)

	app.Get("/users", GetUsers(svc))
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

	svc.On(
		"GetUser",
		mock.Anything,
		mock.Anything,
	).
		Return(
			&dto.UserDetail{
				ID:     1,
				Mobile: "0912",
				Email:  "test@test.com",
			},
			nil,
		)

	app.Get("/users/:id", GetUser(svc))

	req := httptest.NewRequest("GET", "/users/1", nil)
	resp, _ := app.Test(req, fiber.TestConfig{Timeout: 0})

	assert.Equal(t, 200, resp.StatusCode)
}
