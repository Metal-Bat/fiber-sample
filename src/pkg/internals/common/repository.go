package common

import (
	"sample/src/initializers"

	"github.com/gofiber/fiber/v3"
)

type CommonRepository interface {
	Teapot(c fiber.Ctx) (string, error)
}

type repository struct {
}

func NewCommonRepository() CommonRepository {
	return &repository{}
}

func (r *repository) Teapot(c fiber.Ctx) (string, error) {
	_, span := initializers.Tracer.Start(c.Context(), "repository.Teapot")
	defer span.End()

	return "i am a teapot", nil
}
