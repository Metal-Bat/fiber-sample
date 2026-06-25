package api

import (
	"context"

	"github.com/gofiber/fiber/v3"

	"sample/src/api/routes"
	"sample/src/initializers"
	"sample/src/pkg/base"
	"sample/src/pkg/entities"
	"sample/src/pkg/internals/common"
	"sample/src/pkg/internals/user"
	"sample/src/pkg/mapper"
	"sample/src/pkg/utils"
)

func InitialApi(app *fiber.App) {
	apiV1 := app.Group("/api/v1")

	commonApi := apiV1.Group("/common")
	commonRepository := common.NewCommonRepository()
	commonService := common.NewCommonService(commonRepository)
	routes.CommonRouter(commonApi, commonService)

	userApi := apiV1.Group("/user")
	userMapper := &mapper.UserConverterImpl{}
	userHooks := base.Hooks[entities.User]{
		BeforeCreate: func(ctx context.Context, u *entities.User) error {
			hashed, err := utils.HashPassword(ctx, u.Password)
			if err != nil {
				return err
			}
			u.Password = hashed
			return nil
		},
		BeforeUpdate: func(ctx context.Context, u *entities.User) error {
			if u.Password != "" {
				hashed, err := utils.HashPassword(ctx, u.Password)
				if err != nil {
					return err
				}
				u.Password = hashed
			}
			return nil
		},
	}
	userBaseRepo := base.NewRepository[entities.User](initializers.DB, "User", "Permissions")
	userBaseSvc := base.NewService(
		userBaseRepo,
		userMapper,
		userHooks,
		"User",
	)
	userRepo := user.NewUserRepository()
	userService := user.NewUserService(userBaseSvc, userRepo, userMapper)
	routes.UserRouter(userApi, userService)
}
