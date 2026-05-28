package api

import (
	"sample/src/api/routes"
	"sample/src/pkg/internals/common"
	"sample/src/pkg/internals/user"
	"sample/src/pkg/mapper"

	"github.com/gofiber/fiber/v3"
)

func InitialApi(app *fiber.App) {
	apiV1 := app.Group("/api/v1")

	// common
	commonApi := apiV1.Group("/common")
	commonRepository := common.NewCommonRepository()
	commonService := common.NewCommonService(commonRepository)
	routes.CommonRouter(commonApi, commonService)

	// user
	userApi := apiV1.Group("/user")
	userRepository := user.NewUserRepository()
	userMapper := mapper.NewUserMapper()
	userService := user.NewUserService(userRepository, userMapper)
	routes.UserRouter(userApi, userService)

}
