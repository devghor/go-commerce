package handlers

import (
	"go-commerce/internal/api/rest"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	handler := UserHandler{}
	app.Get("/register", handler.Register)
	app.Get("/login", handler.Login)
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User registered successfully",
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User logged in successfully",
	})
}
