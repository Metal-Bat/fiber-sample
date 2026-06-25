package routes

import (
	"github.com/gofiber/fiber/v3"

	"sample/src/api/handler"
	"sample/src/api/middleware"
	"sample/src/pkg/base"
	"sample/src/pkg/dto"
	"sample/src/pkg/internals/user"
)

var userFilterable = []string{"id", "mobile", "national_code", "created_at", "updated_at"}

func UserRouter(app fiber.Router, service user.Service) {
	crud := base.Service[dto.UserInfo, dto.UserDetail, dto.CreateUser, dto.UpdateUser](service)

	app.Get(
		"",
		middleware.Protected(),
		middleware.RequirePermission(middleware.ADMIN),
		handler.GetAll(crud, userFilterable),
	)

	app.Get(
		"/:id<int>",
		middleware.Protected(),
		middleware.RequirePermission(middleware.ADMIN),
		handler.GetOne(crud),
	)

	app.Post(
		"",
		middleware.Protected(),
		middleware.RequirePermission(middleware.ADMIN),
		handler.Create(crud),
	)

	app.Put(
		"/:id<int>",
		middleware.Protected(),
		middleware.RequirePermission(middleware.ADMIN),
		handler.Update(crud),
	)

	app.Delete(
		"/:id<int>",
		middleware.Protected(),
		middleware.RequirePermission(middleware.ADMIN),
		handler.Delete(crud),
	)

	app.Post("/login", handler.Login(service))
}
