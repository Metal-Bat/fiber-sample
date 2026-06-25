package handler

import (
	"errors"

	"github.com/gofiber/fiber/v3"

	"sample/src/api/presenter"
	"sample/src/initializers"
	"sample/src/pkg/internals/user"
	"sample/src/pkg/utils/customErrors"
)

func Login(service user.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		_, span := initializers.Tracer.Start(c.Context(), "handler.User.Login")
		defer span.End()

		var input presenter.LoginInput
		if err := c.Bind().Body(&input); err != nil {
			span.RecordError(err)
			return presenter.InputSerializerError(c)
		}

		result, err := service.Login(c.Context(), input.Mobile, input.Password)
		if err != nil {
			var Err *customErrors.InternalError
			if errors.As(err, &Err) {
				return presenter.ReturnInternalError(c, Err)
			}
			return presenter.InvalidUserPassError(c)
		}

		return presenter.SuccessfulResponse(c, result)
	}
}
