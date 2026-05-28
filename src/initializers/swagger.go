package initializers

import (
	"github.com/gofiber/contrib/v3/swaggerui"
)

var SwaggerConfig *swaggerui.Config

func SetUpSwagger() {
	SwaggerConfig = &swaggerui.Config{
		FilePath: "src/api/docs/openapi.yaml",
		Path:     "/swagger-ui",
	}
}
