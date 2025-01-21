package handlers

import (
	"clean-architecture/internal/application/usecases"
	"clean-architecture/internal/domain/entities"

	"github.com/gofiber/fiber/v2"
)

type SubjectHandler struct {
	subjectUsecase usecases.SubjectUsecase
}

func NewSubjectHandler(su usecases.SubjectUsecase) *SubjectHandler {
	return &SubjectHandler{
		subjectUsecase: su,
	}
}

// POST /api/subjects
func (h *SubjectHandler) CreateSubject(c *fiber.Ctx) error {
	var subject entities.Subject
	if err := c.BodyParser(&subject); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.subjectUsecase.CreateSubject(&subject); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(subject)
}

// GET /api/subjects/:id
func (h *SubjectHandler) GetSubjectByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	subject, err := h.subjectUsecase.GetSubjectByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if subject == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Subject not found"})
	}

	return c.JSON(subject)
}

// PUT /api/subjects/:id
func (h *SubjectHandler) UpdateSubject(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var subject entities.Subject
	if err := c.BodyParser(&subject); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	subject.SubjectID = id

	if err := h.subjectUsecase.UpdateSubject(&subject); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(subject)
}

// DELETE /api/subjects/:id
func (h *SubjectHandler) DeleteSubject(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.subjectUsecase.DeleteSubject(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Subject deleted successfully"})
}

// GET /api/subjects
func (h *SubjectHandler) GetAllSubjects(c *fiber.Ctx) error {
	subjects, err := h.subjectUsecase.GetAllSubjects()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(subjects)
}
