package utils

import (
	"context"
	"os"
	"sample/src/initializers"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJwtToken(ctx context.Context, mobile string, permissions []string, expires_at int64) (string, error) {
	_, span := initializers.Tracer.Start(ctx, "utils.CreateJwtToken")
	defer span.End()

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["mobile"] = mobile
	claims["expires_at"] = expires_at
	claims["permissions"] = permissions

	t, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		span.RecordError(err)
		return "", err
	}
	return t, nil
}
