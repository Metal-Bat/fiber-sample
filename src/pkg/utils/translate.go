package utils

import (
	"sample/src/initializers"

	contribI18nMiddleware "github.com/gofiber/contrib/v3/i18n"
	"github.com/gofiber/fiber/v3"
)

func TranslateMessage(c fiber.Ctx, message string) (string, error) {
	_, span := initializers.Tracer.Start(c.Context(), "utils.TranslateMessage")
	defer span.End()

	translatedMessage, err := contribI18nMiddleware.Localize(c, message)
	if err != nil {
		span.RecordError(err)
		return message, nil
	}
	return translatedMessage, nil
}
