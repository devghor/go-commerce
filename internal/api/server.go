package api

import (
	"go-commerce/configs"
	"go-commerce/internal/api/rest"
	"go-commerce/internal/api/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config configs.AppConfig) {

	app := fiber.New()

	rh := &rest.RestHandler{
		App: app,
	}

	setupRoutes(rh)

	app.Listen(":" + config.Port)
}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoutes(rh)
}
