package utils

import (
	"sample/src/initializers"

	"github.com/gofiber/fiber/v3"
)

func TranslateMessage(c fiber.Ctx, message string) (string, error) {
	_, span := initializers.Tracer.Start(c.Context(), "utils.TranslateMessage")
	defer span.End()

	translatedMessage, err := initializers.Translator.Localize(c, message)
	if err != nil {
		span.RecordError(err)
		return message, nil
	}
	return translatedMessage, nil
}
