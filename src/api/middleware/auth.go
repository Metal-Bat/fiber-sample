package middleware

import (
	"os"
	"sample/src/api/presenter"
	"sample/src/initializers"

	jwtMiddleware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
)

func Protected() func(fiber.Ctx) error {
	return jwtMiddleware.New(jwtMiddleware.Config{
		SigningKey:   jwtMiddleware.SigningKey{Key: []byte(os.Getenv("SECRET"))},
		ErrorHandler: jwtError,
	})
}

func jwtError(c fiber.Ctx, err error) error {
	_, span := initializers.Tracer.Start(c.Context(), "handler.GetTeapot")
	defer span.End()
	span.RecordError(err)
	return presenter.AuthError(c)
}
