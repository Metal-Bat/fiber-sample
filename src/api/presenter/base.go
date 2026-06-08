package presenter

import (
	"sample/src/initializers"
	"sample/src/pkg/dto"
	"sample/src/pkg/utils"

	"github.com/gofiber/fiber/v3"
)

func ReturnSimpleSuccess(c fiber.Ctx) error {
	_, span := initializers.Tracer.Start(c.Context(), "presenter.ReturnSimpleSuccess")
	defer span.End()

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"result":  nil,
			"error":   nil,
			"success": true,
			"code":    fiber.StatusOK,
		},
	)
}

func ReturnSuccessWithMessage(c fiber.Ctx, message string) error {
	_, span := initializers.Tracer.Start(c.Context(), "presenter.ReturnSuccessWithMessage")
	defer span.End()

	translatedMessage, _ := utils.TranslateMessage(c, message)
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"result":  translatedMessage,
			"error":   nil,
			"success": true,
			"code":    fiber.StatusOK,
		},
	)
}

func ReturnSimpleCreateSuccess(c fiber.Ctx) error {
	_, span := initializers.Tracer.Start(c.Context(), "presenter.ReturnSimpleCreateSuccess")
	defer span.End()

	return c.Status(fiber.StatusCreated).JSON(
		fiber.Map{
			"result":  nil,
			"error":   nil,
			"success": true,
			"code":    fiber.StatusCreated,
		},
	)
}

func ReturnError(c fiber.Ctx, errorMessage *fiber.Error) error {
	_, span := initializers.Tracer.Start(c.Context(), "presenter.ReturnError")
	defer span.End()

	ReturnErrorMessage, _ := utils.TranslateMessage(c, errorMessage.Message)
	return c.Status(fiber.StatusBadRequest).JSON(
		fiber.Map{
			"result":  nil,
			"error":   ReturnErrorMessage,
			"success": false,
			"code":    fiber.StatusBadRequest,
		},
	)
}

func AuthError(c fiber.Ctx) error {
	_, span := initializers.Tracer.Start(c.Context(), "presenter.AuthError")
	defer span.End()

	AuthErrorMessage, _ := utils.TranslateMessage(c, "invalid or expired token")
	return c.Status(fiber.StatusUnauthorized).JSON(
		fiber.Map{
			"result":  nil,
			"error":   AuthErrorMessage,
			"success": false,
			"code":    fiber.StatusUnauthorized,
		},
	)
}

func InvalidUserPassError(c fiber.Ctx) error {
	_, span := initializers.Tracer.Start(c.Context(), "presenter.InvalidUserPassError")
	defer span.End()

	InvalidUserPassMessage, _ := utils.TranslateMessage(c, "invalid user or password")
	return c.Status(fiber.StatusUnauthorized).JSON(
		fiber.Map{
			"result":  nil,
			"error":   InvalidUserPassMessage,
			"success": false,
			"code":    fiber.StatusUnauthorized,
		},
	)
}

func ForbiddenError(c fiber.Ctx) error {
	_, span := initializers.Tracer.Start(c.Context(), "presenter.ForbiddenError")
	defer span.End()

	ForbiddenMessage, _ := utils.TranslateMessage(c, "permission denied")
	return c.Status(fiber.StatusForbidden).JSON(
		fiber.Map{
			"result":  nil,
			"error":   ForbiddenMessage,
			"success": false,
			"code":    fiber.StatusForbidden,
		},
	)
}

func ReturnNotFound(c fiber.Ctx) error {
	_, span := initializers.Tracer.Start(c.Context(), "presenter.ReturnNotFound")
	defer span.End()

	ReturnNotFoundMessage, _ := utils.TranslateMessage(c, "entity not found")
	return c.Status(fiber.StatusNotFound).JSON(
		fiber.Map{
			"result":  nil,
			"error":   ReturnNotFoundMessage,
			"success": false,
			"code":    fiber.StatusNotFound,
		},
	)
}

func InputSerializerError(c fiber.Ctx) error {
	_, span := initializers.Tracer.Start(c.Context(), "presenter.InputSerializerError")
	defer span.End()

	InputSerializerErrorMessage, _ := utils.TranslateMessage(c, "unprocessable input")
	return c.Status(fiber.StatusUnprocessableEntity).JSON(
		fiber.Map{
			"result":  nil,
			"error":   InputSerializerErrorMessage,
			"success": false,
			"code":    fiber.StatusUnprocessableEntity,
		},
	)

}

func SuccessfulResponse[T any](c fiber.Ctx, item T) error {
	_, span := initializers.Tracer.Start(c.Context(), "presenter.SuccessfulResponse")
	defer span.End()

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"result":  item,
			"error":   nil,
			"success": true,
			"code":    fiber.StatusOK,
		},
	)
}

func SuccessfulCreateResponse[T any](c fiber.Ctx, item T) error {
	_, span := initializers.Tracer.Start(c.Context(), "presenter.SuccessfulCreateResponse")
	defer span.End()

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"result":  item,
			"error":   nil,
			"success": true,
			"code":    fiber.StatusCreated,
		},
	)
}

func SuccessfulUpdateResponse(c fiber.Ctx) error {
	_, span := initializers.Tracer.Start(c.Context(), "presenter.SuccessfulUpdateResponse")
	defer span.End()

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"result":  nil,
			"error":   nil,
			"success": true,
			"code":    fiber.StatusNoContent,
		},
	)
}

func SuccessfulDeleteResponse(c fiber.Ctx) error {
	_, span := initializers.Tracer.Start(c.Context(), "presenter.SuccessfulDeleteResponse")
	defer span.End()

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"result":  nil,
			"error":   nil,
			"success": true,
			"code":    fiber.StatusNoContent,
		},
	)
}

func SuccessfulPaginatedResponse[T any](
	c fiber.Ctx,
	page dto.PaginationStructure,
	items T,
	count int64,
) error {
	_, span := initializers.Tracer.Start(c.Context(), "presenter.SuccessfulPaginatedResponse")
	defer span.End()

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"result": fiber.Map{
				"page":  page.Page,
				"size":  page.Size,
				"count": count,
				"items": items,
			},
			"error":   nil,
			"success": true,
			"code":    fiber.StatusOK,
		},
	)
}
