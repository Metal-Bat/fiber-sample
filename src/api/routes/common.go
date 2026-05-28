package routes

import (
	"sample/src/api/handler"
	"sample/src/api/middleware"
	"sample/src/pkg/internals/common"

	"github.com/gofiber/fiber/v3"
)

func CommonRouter(app fiber.Router, service common.Service) {
	app.Get(
		"/teapot",
		middleware.Protected(),
		middleware.RequirePermission(middleware.TEAPOT),
		handler.GetTeapot(service),
	)
}
