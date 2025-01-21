package routes

import (
	"clean-architecture/config"
	"clean-architecture/internal/infrastructure/http/handlers"
	"clean-architecture/internal/infrastructure/http/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App,
	cfg *config.Config,

	// Auth
	authHandler *handlers.AuthHandler,

	// Student
	studentHandler *handlers.StudentHandler,

	// Teacher
	teacherHandler *handlers.TeacherHandler,

	// Subject
	subjectHandler *handlers.SubjectHandler,

	// User
	userHandler *handlers.UserHandler,
) {

	// -------------------------------
	// Public routes (no JWT required)
	// -------------------------------
	authGroup := app.Group("/auth")
	authGroup.Post("/register", authHandler.Register)
	authGroup.Post("/login", authHandler.Login)
	authGroup.Post("/refresh", authHandler.RefreshToken)

	// -------------------------------
	// Protected routes (JWT required)
	// -------------------------------
	api := app.Group("/api", middlewares.JWTProtected(cfg))

	// Auth
	api.Post("/logout", authHandler.Logout)

	// Students
	api.Post("/students", studentHandler.CreateStudent)
	api.Get("/students/:id", studentHandler.GetStudent)
	api.Put("/students/:id", studentHandler.UpdateStudent)
	api.Delete("/students/:id", studentHandler.DeleteStudent)
	api.Get("/students", studentHandler.GetAllStudents)

	// ตัวอย่างอัปเดตหลายตาราง
	api.Post("/students/:id/assign-subjects", studentHandler.AssignStudentSubjects)

	// Teachers
	api.Post("/teachers", teacherHandler.CreateTeacher)
	api.Get("/teachers/:id", teacherHandler.GetTeacherByID)
	api.Put("/teachers/:id", teacherHandler.UpdateTeacher)
	api.Delete("/teachers/:id", teacherHandler.DeleteTeacher)
	api.Get("/teachers", teacherHandler.GetAllTeachers)
	api.Post("/teachers/:id/assign-subjects", teacherHandler.AssignTeacherSubjects)

	// Subjects
	api.Post("/subjects", subjectHandler.CreateSubject)
	api.Get("/subjects/:id", subjectHandler.GetSubjectByID)
	api.Put("/subjects/:id", subjectHandler.UpdateSubject)
	api.Delete("/subjects/:id", subjectHandler.DeleteSubject)
	api.Get("/subjects", subjectHandler.GetAllSubjects)

	// Users
	api.Get("/users", userHandler.GetAllUsers)
	api.Get("/users/:id", userHandler.GetUserByID)
}
