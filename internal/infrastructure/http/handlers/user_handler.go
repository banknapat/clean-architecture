package handlers

import (
	"github.com/gofiber/fiber/v2"

	"clean-architecture/internal/application/usecases"
)

type UserHandler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHandler(uu usecases.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: uu}
}

// GET /api/users/:id
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	user, err := h.userUsecase.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

// GET /api/users
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userUsecase.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}
