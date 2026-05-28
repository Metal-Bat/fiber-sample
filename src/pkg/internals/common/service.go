package common

import (
	"sample/src/initializers"

	"github.com/gofiber/fiber/v3"
)

type Service interface {
	Teapot(c fiber.Ctx) (string, error)
}

type service struct {
	repository CommonRepository
}

func NewCommonService(r CommonRepository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Teapot(c fiber.Ctx) (string, error) {
	_, span := initializers.Tracer.Start(c.Context(), "service.Teapot")
	defer span.End()

	return s.repository.Teapot(c)
}
