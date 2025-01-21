package handlers

import (
	"clean-architecture/internal/application/usecases"
	"clean-architecture/internal/domain/entities"

	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	studentUsecase usecases.StudentUsecase
}

func NewStudentHandler(su usecases.StudentUsecase) *StudentHandler {
	return &StudentHandler{
		studentUsecase: su,
	}
}

// POST /api/students
func (h *StudentHandler) CreateStudent(c *fiber.Ctx) error {
	var student entities.Student
	if err := c.BodyParser(&student); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.studentUsecase.CreateStudent(&student)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(student)
}

// GET /api/students/:id
func (h *StudentHandler) GetStudent(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	student, err := h.studentUsecase.GetStudentByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if student == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Student not found"})
	}

	return c.JSON(student)
}

// PUT /api/students/:id
func (h *StudentHandler) UpdateStudent(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var student entities.Student
	if err := c.BodyParser(&student); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	student.StudentID = id

	err = h.studentUsecase.UpdateStudent(&student)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(student)
}

// DELETE /api/students/:id
func (h *StudentHandler) DeleteStudent(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	err = h.studentUsecase.DeleteStudent(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Student deleted successfully"})
}

// GET /api/students
func (h *StudentHandler) GetAllStudents(c *fiber.Ctx) error {
	students, err := h.studentUsecase.GetAllStudents()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(students)
}

// ตัวอย่างอัปเดตหลายตารางพร้อมกัน
// POST /api/students/:id/assign-subjects
func (h *StudentHandler) AssignStudentSubjects(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var request struct {
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		SubjectIDs []int  `json:"subject_ids"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = h.studentUsecase.UpdateStudentAndAssignSubjects(id, request.FirstName, request.LastName, request.SubjectIDs)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Updated student and assigned subjects successfully"})
}
