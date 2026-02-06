package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nezts08/discordOAuth2WithGo/controller"
)

func SetupAuthRoutes(app *fiber.App) {
	authGroup := app.Group("auth")

	authGroup.Get("/login", func(c fiber.Ctx) error {
		return controller.AuthLoginController(c)
	})

	authGroup.Get("/redirect", func(c fiber.Ctx) error {
		return controller.AuthRedirectController(c)
	})

	authGroup.Get("/logout", func(c fiber.Ctx) error {
		return controller.AuthLogoutController(c)
	})

}
