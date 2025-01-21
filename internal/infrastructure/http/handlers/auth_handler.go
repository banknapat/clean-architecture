package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"clean-architecture/internal/application/usecases"
	"clean-architecture/internal/domain/value_objects"
)

type AuthHandler struct {
	authUsecase usecases.AuthUsecase
}

func NewAuthHandler(au usecases.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: au,
	}
}

// POST /auth/register
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var creds value_objects.Credentials
	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.authUsecase.Register(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "User registered successfully"})
}

// POST /auth/login
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var creds value_objects.Credentials
	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	tokens, err := h.authUsecase.Login(&creds)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	// set cookie
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    tokens.AccessToken,
		Expires:  time.Now().Add(time.Hour),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}

// POST /auth/logout
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	if userID == nil {
		// หาก library jwt v3 ของ fiber ใช้ "user_id" หรือ "user" หรือ "claims" => ปรับตามจริง
		// ตัวอย่างอาจต้อง parse claims ก่อนก็ได้
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// สมมติว่า user_id เป็น int
	err := h.authUsecase.Logout(userID.(int))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// clear cookie
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{"message": "Logged out successfully"})
}

// POST /auth/refresh
func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	var body struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	tokens, err := h.authUsecase.RefreshToken(body.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    tokens.AccessToken,
		Expires:  time.Now().Add(time.Hour),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}
