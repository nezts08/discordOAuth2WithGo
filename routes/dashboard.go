package routes

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nezts08/discordOAuth2WithGo/auth"
	"github.com/nezts08/discordOAuth2WithGo/repository"
)

func SetupDashboardRoutes(app *fiber.App) {
	dashboardGroup := app.Group("dashboard")

	dashboardGroup.Get("/", auth.IsAuthorized, func(c fiber.Ctx) error {
		claims := c.Locals("user").(jwt.MapClaims)

		userID, ok := claims["user_id"].(string)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid user_id")
		}

		existingUser, err := repository.FindUserByDiscordID(userID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error fetching user from database")
		}

		var ClientID string = os.Getenv("CLIENT_ID")

		return c.Render("dashboard", fiber.Map{
			"Title":    "Dashboard",
			"Username": existingUser.Username,
			"Guilds":   existingUser.Guilds,
			"ClientID": ClientID,
		})
	})
}
