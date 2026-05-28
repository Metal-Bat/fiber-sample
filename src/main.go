package main

import (
	"context"

	"sample/src/api"
	"sample/src/initializers"

	"fmt"
	"log"
	"os"

	"github.com/gofiber/contrib/v3/fgprof"
	"github.com/gofiber/contrib/v3/monitor"
	"github.com/gofiber/contrib/v3/otel"
	"github.com/gofiber/contrib/v3/swaggerui"
	"github.com/gofiber/fiber/v3"

	"github.com/gofiber/contrib/v3/i18n"
	fiberLogger "github.com/gofiber/contrib/v3/zap"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func init() {
	initializers.SetUpDb()
	initializers.SyncDatabase()
	initializers.SetUpRedis()
	initializers.SetUpTranslator()
	initializers.SetUpCors()
	initializers.SetUpJwt()
	initializers.SetUpSwagger()
}

func main() {

	tp := initializers.InitTracer()
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	logger := initializers.InitLogger()
	defer func() { _ = logger.Sync() }()

	app := fiber.New()
	app.Use(fgprof.New())
	app.Use(otel.Middleware())
	app.Use(recover.New())
	app.Use(i18n.New(initializers.TranslateConfig))
	app.Use(cors.New(*initializers.CorsConfig))
	app.Use(swaggerui.New(*initializers.SwaggerConfig))
	app.Get("/api/v1/metrics", monitor.New())
	app.Use(
		fiberLogger.New(
			fiberLogger.Config{
				Logger: logger,
				Fields: initializers.LoggerFields(),
			},
		),
	)

	api.InitialApi(app)

	server := fmt.Sprintf("%s:%s", os.Getenv("SERVER"), os.Getenv("PORT"))
	if err := app.Listen(server); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
