package middlewares

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/gofiber/jwt/v3"

	"clean-architecture/config"
)

func JWTProtected(cfg *config.Config) fiber.Handler {
	return jwt.New(jwt.Config{
		SigningKey:   []byte(cfg.JWTSecret),
		ErrorHandler: jwtError,
		TokenLookup:  "header:Authorization,cookie:jwt",
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": errors.New("Unauthorized or invalid token").Error(),
		})
	}
	return nil
}
