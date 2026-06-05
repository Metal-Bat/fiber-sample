package middleware

import (
	"errors"
	"sample/src/api/presenter"
	"sample/src/initializers"
	"slices"

	jwtMiddleware "github.com/gofiber/contrib/v3/jwt"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gofiber/fiber/v3"
)

const (
	TEAPOT = "teapot"
	ADMIN  = "admin"
)

var AllPermissions = []string{
	TEAPOT,
	ADMIN,
}

func getPermissions(c fiber.Ctx) ([]string, error) {
	_, span := initializers.Tracer.Start(c.Context(), "middleware.getPermissions")
	defer span.End()

	token := jwtMiddleware.FromContext(c)
	if token == nil {
		return nil, errors.New("token not found in context")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid jwt claims")
	}

	rawPerms, ok := claims["permissions"]
	if !ok {
		return nil, errors.New("permissions not found in token")
	}

	permissionInterface, ok := rawPerms.([]interface{})
	if !ok {
		return nil, errors.New("invalid permissions format")
	}

	perms := make([]string, len(permissionInterface))
	for i, p := range permissionInterface {
		perms[i], _ = p.(string)
	}

	return perms, nil
}

func RequirePermission(required string) fiber.Handler {
	return func(c fiber.Ctx) error {
		_, span := initializers.Tracer.Start(c.Context(), "middleware.RequirePermission")
		defer span.End()

		perms, err := getPermissions(c)
		if err != nil {
			span.RecordError(err)
			return presenter.AuthError(c)
		}

		if slices.Contains(perms, required) {
			return c.Next()
		}

		return presenter.ForbiddenError(c)
	}
}
