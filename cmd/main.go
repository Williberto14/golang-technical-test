package main

import (
	"golang-technical-test/config"
	"golang-technical-test/database"
	"golang-technical-test/internal/delivery/http"
	"golang-technical-test/internal/repository"
	"golang-technical-test/internal/usecase"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Load the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err.Error())
	}

	// Initialize the database
	db, err := database.NewDatabase(cfg.DB)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Initialize the repositories
	studentRepo := repository.NewStudentRepository(db)
	courseRepo := repository.NewCourseRepository(db)
	professorRepo := repository.NewProfessorRepository(db)
	gradeRepo := repository.NewGradeRepository(db)
	enrollmentRepo := repository.NewEnrollmentRepository(db)

	// Initialize the usecases
	studentUsecase := usecase.NewStudentUsecase(studentRepo)
	courseUsecase := usecase.NewCourseUsecase(courseRepo)
	professorUsecase := usecase.NewProfessorUsecase(professorRepo)
	gradeUsecase := usecase.NewGradeUsecase(gradeRepo)
	enrollmentUsecase := usecase.NewEnrollmentUsecase(enrollmentRepo)

	// Initialize the router
	router := gin.Default()

	// Initialize the handlers
	http.NewLoginHandler(router)
	http.NewStudentHandler(studentUsecase, router)
	http.NewCourseHandler(courseUsecase, router)
	http.NewProfessorHandler(professorUsecase, router)
	http.NewGradeHandler(gradeUsecase, router)
	http.NewEnrollmentHandler(enrollmentUsecase, router)

	// Run the server
	router.Run(":7777")
}
