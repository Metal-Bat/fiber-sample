package handler

import (
	"sample/src/api/presenter"
	"sample/src/initializers"
	"sample/src/pkg/internals/common"

	"github.com/gofiber/fiber/v3"
)

func GetTeapot(service common.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		_, span := initializers.Tracer.Start(c.Context(), "handler.GetTeapot")
		defer span.End()

		text, err := service.Teapot(c)

		if err != nil {
			span.RecordError(err)
		}
		return presenter.ReturnSuccessWithMessage(c, text)
	}
}
