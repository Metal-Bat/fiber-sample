package mapper

import (
	"sample/src/initializers"

	"github.com/gofiber/fiber/v3"
)

func MapSlice[A any, B any](
	c fiber.Ctx,
	in []*A,
	fn func(fiber.Ctx, *A) *B,
) []*B {
	_, span := initializers.Tracer.Start(c.Context(), "base.MapSlice")
	defer span.End()

	out := make([]*B, len(in))
	for i, v := range in {
		out[i] = fn(c, v)
	}
	return out
}
