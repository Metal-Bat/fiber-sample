package utils

import (
	"context"
	"sample/src/initializers"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(ctx context.Context, password string) (string, error) {
	_, span := initializers.Tracer.Start(ctx, "utils.HashPassword")
	defer span.End()

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		span.RecordError(err)
		return "", err
	}
	return string(hashed), nil
}

func ComparHash(ctx context.Context, hashed string, password string) error {
	_, span := initializers.Tracer.Start(ctx, "utils.ComparHash")
	defer span.End()

	return bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte(password),
	)
}
