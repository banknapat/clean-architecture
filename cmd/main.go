package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"clean-architecture/config"
	"clean-architecture/internal/application/usecases"
	"clean-architecture/internal/domain/repositories"
	"clean-architecture/internal/infrastructure/http/handlers"
	"clean-architecture/internal/infrastructure/http/routes"
	postgresRepos "clean-architecture/internal/infrastructure/postgres/repositories"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Connect to DB
	db, err := gorm.Open(postgres.Open(cfg.GetPostgresDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// =====
	// ถ้าคุณต้องการ AutoMigrate:
	// db.AutoMigrate(&entities.User{}, &entities.Student{}, &entities.Teacher{}, &entities.Subject{})
	// แต่ใน production แนะนำให้ใช้ migrations แบบ SQL ด้านบน
	// =====

	// Repositories
	var (
		userRepo    repositories.UserRepository    = postgresRepos.NewUserRepository(db)
		studentRepo repositories.StudentRepository = postgresRepos.NewStudentRepository(db)
		teacherRepo repositories.TeacherRepository = postgresRepos.NewTeacherRepository(db)
		subjectRepo repositories.SubjectRepository = postgresRepos.NewSubjectRepository(db)
	)

	// Usecases
	var (
		authUsecase    usecases.AuthUsecase    = usecases.NewAuthUsecase(userRepo, cfg)
		studentUsecase usecases.StudentUsecase = usecases.NewStudentUsecase(studentRepo)
		teacherUsecase usecases.TeacherUsecase = usecases.NewTeacherUsecase(teacherRepo)
		subjectUsecase usecases.SubjectUsecase = usecases.NewSubjectUsecase(subjectRepo)
		userUsecase    usecases.UserUsecase    = usecases.NewUserUsecase(userRepo)
	)

	// Handlers
	authHandler := handlers.NewAuthHandler(authUsecase)
	studentHandler := handlers.NewStudentHandler(studentUsecase)
	teacherHandler := handlers.NewTeacherHandler(teacherUsecase)
	subjectHandler := handlers.NewSubjectHandler(subjectUsecase)
	userHandler := handlers.NewUserHandler(userUsecase)

	// Fiber
	app := fiber.New()
	app.Use(logger.New()) // middleware logger

	// Setup Routes
	routes.SetupRoutes(
		app, cfg,
		authHandler,
		studentHandler,
		teacherHandler,
		subjectHandler,
		userHandler,
	)

	// Start Server
	log.Println("Starting server on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
