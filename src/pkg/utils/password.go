package utils

import (
	"sample/src/initializers"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(c fiber.Ctx, password string) (string, error) {
	_, span := initializers.Tracer.Start(c.Context(), "utils.HashPassword")
	defer span.End()

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		span.RecordError(err)
		return "", err
	}
	return string(hashed), nil
}

func ComparHash(c fiber.Ctx, hashed string, password string) error {
	_, span := initializers.Tracer.Start(c.Context(), "utils.ComparHash")
	defer span.End()

	return bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte(password),
	)
}
