package handler

import (
	"errors"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCommonService struct {
	mock.Mock
}

func (m *mockCommonService) Teapot(c fiber.Ctx) (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func TestGetTeapot_Success(t *testing.T) {
	app := fiber.New()
	mockSvc := new(mockCommonService)
	mockSvc.
		On("Teapot", mock.Anything).
		Return("ok teapot", nil)

	app.Get("/teapot", GetTeapot(mockSvc))

	req := httptest.NewRequest("GET", "/teapot", nil)

	resp, err := app.Test(req, fiber.TestConfig{
		Timeout: 0,
	})

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	assert.Contains(t, string(body), "ok teapot")

}

func TestGetTeapot_Error(t *testing.T) {
	app := fiber.New()

	mockSvc := new(mockCommonService)
	mockSvc.
		On("Teapot", mock.Anything).
		Return("", errors.New("boom"))

	app.Get("/teapot", GetTeapot(mockSvc))

	req := httptest.NewRequest("GET", "/teapot", nil)

	resp, err := app.Test(req, fiber.TestConfig{
		Timeout: 0,
	})

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
