package routes

import (
	"sample/src/api/handler"
	"sample/src/api/middleware"
	"sample/src/pkg/internals/user"

	"github.com/gofiber/fiber/v3"
)

func UserRouter(app fiber.Router, service user.Service) {
	// user crud
	app.Get(
		"",
		middleware.Protected(),
		middleware.RequirePermission(middleware.ADMIN),
		handler.GetUsers(service),
	)
	app.Get(
		"/:id<int>",
		middleware.Protected(),
		middleware.RequirePermission(middleware.ADMIN),
		handler.GetUser(service),
	)
	app.Post(
		"",
		middleware.Protected(),
		middleware.RequirePermission(middleware.ADMIN),
		handler.CreateUser(service),
	)
	app.Put(
		"/:id<int>",
		middleware.Protected(),
		middleware.RequirePermission(middleware.ADMIN),
		handler.UpdateUser(service),
	)
	app.Delete(
		"/:id<int>",
		middleware.Protected(),
		middleware.RequirePermission(middleware.ADMIN),
		handler.DeleteUser(service),
	)

	// services
	app.Post("/login", handler.Login(service))
}
