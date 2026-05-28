package initializers

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v3/middleware/cors"
)

var CorsConfig *cors.Config

func SetUpCors() {
	CorsConfig = &cors.Config{
		AllowOrigins: strings.Split(os.Getenv("ALLOWED_ORIGINS"), ","),
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}
}
