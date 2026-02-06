package auth

import (
	"github.com/gofiber/fiber/v3"
)

func IsAuthorized(c fiber.Ctx) error {
	token := c.Cookies("auth")
	if token == "" {
		return c.Redirect().To("/")
	}

	claims, err := ParseJwt(token)

	if err != nil {
		return c.Redirect().To("/")
	}

	c.Locals("user", claims)
	return c.Next()
}
