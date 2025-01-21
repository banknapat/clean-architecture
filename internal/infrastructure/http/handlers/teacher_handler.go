package handlers

import (
	"clean-architecture/internal/application/usecases"
	"clean-architecture/internal/domain/entities"

	"github.com/gofiber/fiber/v2"
)

type TeacherHandler struct {
	teacherUsecase usecases.TeacherUsecase
}

func NewTeacherHandler(tu usecases.TeacherUsecase) *TeacherHandler {
	return &TeacherHandler{
		teacherUsecase: tu,
	}
}

// POST /api/teachers
func (h *TeacherHandler) CreateTeacher(c *fiber.Ctx) error {
	var teacher entities.Teacher
	if err := c.BodyParser(&teacher); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.teacherUsecase.CreateTeacher(&teacher); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(teacher)
}

// GET /api/teachers/:id
func (h *TeacherHandler) GetTeacherByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	teacher, terr := h.teacherUsecase.GetTeacherByID(id)
	if terr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": terr.Error()})
	}

	if teacher == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Teacher not found"})
	}

	return c.JSON(teacher)
}

// PUT /api/teachers/:id
func (h *TeacherHandler) UpdateTeacher(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var teacher entities.Teacher
	if err := c.BodyParser(&teacher); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	teacher.TeacherID = id

	if err := h.teacherUsecase.UpdateTeacher(&teacher); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(teacher)
}

// DELETE /api/teachers/:id
func (h *TeacherHandler) DeleteTeacher(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.teacherUsecase.DeleteTeacher(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Teacher deleted successfully"})
}

// GET /api/teachers
func (h *TeacherHandler) GetAllTeachers(c *fiber.Ctx) error {
	teachers, err := h.teacherUsecase.GetAllTeachers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(teachers)
}

// POST /api/teachers/:id/assign-subjects
func (h *TeacherHandler) AssignTeacherSubjects(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var req struct {
		SubjectIDs []int `json:"subject_ids"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.teacherUsecase.AssignSubjectsToTeacher(id, req.SubjectIDs); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Assigned subjects to teacher successfully"})
}
