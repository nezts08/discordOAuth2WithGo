package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/pug/v3"
	"github.com/joho/godotenv"
	"github.com/nezts08/discordOAuth2WithGo/auth"
	database "github.com/nezts08/discordOAuth2WithGo/db"
	"github.com/nezts08/discordOAuth2WithGo/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")

	database.Connect()

	engine := pug.New("./views", ".pug")

	engine.Reload(true)

	engine.AddFunc("loggedIn", func(user interface{}) bool {
		return user != nil
	})

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path}\n",
	}))

	app.Use("/static", static.New("./public"))

	app.Get("/", func(c fiber.Ctx) error {
		token := c.Cookies("auth")

		var user interface{} = nil

		if token != "" {
			claims, err := auth.ParseJwt(token)
			if err == nil {
				user = claims
			}
		}

		return c.Render("home", fiber.Map{
			"Title": "Home",
			"User":  user,
		})
	})

	app.Get("/forbidden", func(req fiber.Req, res fiber.Res) error {
		return res.Status(fiber.StatusForbidden).SendString("Forbidden")
	})

	routes.SetupAuthRoutes(app)
	routes.SetupDashboardRoutes(app)

	app.Listen("localhost:"+port, fiber.ListenConfig{
		CertFile:    "./localhost+2.pem",
		CertKeyFile: "./localhost+2-key.pem",
	})

}
