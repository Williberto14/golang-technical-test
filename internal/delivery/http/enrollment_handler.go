package http

import (
	"golang-technical-test/internal/domain"
	"golang-technical-test/internal/usecase"
	"golang-technical-test/middlewares"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type EnrollmentHandler struct {
	EnrollmentUsecase usecase.IEnrollmentUsecase
	path              string
}

var (
	enrollmentHandlerInstance *EnrollmentHandler
	enrollmentHandlerOnce     sync.Once
)

func NewEnrollmentHandler(enrollmentUsecase usecase.IEnrollmentUsecase, router *gin.Engine) *EnrollmentHandler {
	enrollmentHandlerOnce.Do(func() {
		enrollmentHandlerInstance = &EnrollmentHandler{
			EnrollmentUsecase: enrollmentUsecase,
			path:              "/enrollments",
		}
		enrollmentHandlerInstance.setupRoutes(router)
	})
	return enrollmentHandlerInstance
}

func (h *EnrollmentHandler) setupRoutes(router *gin.Engine) {
	JWTGroup := router.Group("/private")
	JWTGroup.Use(middlewares.JWTAuthMiddleware())

	JWTGroup.GET(h.path, h.GetAll)
	JWTGroup.GET(h.path+"/:id", h.GetByID)
	JWTGroup.POST(h.path+"/create", h.Create)
	JWTGroup.PUT(h.path+"/update/:id", h.Update)
	JWTGroup.DELETE(h.path+"/delete/:id", h.Delete)
	JWTGroup.GET(h.path+"/student/:studentID", h.GetByStudentID)
	JWTGroup.GET(h.path+"/course/:courseID", h.GetByCourseID)
}

func (h *EnrollmentHandler) GetAll(c *gin.Context) {
	enrollments, err := h.EnrollmentUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(enrollments) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No enrollments found"})
		return
	}

	c.JSON(http.StatusOK, enrollments)
}

func (h *EnrollmentHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	enrollment, err := h.EnrollmentUsecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, enrollment)
}

func (h *EnrollmentHandler) Create(c *gin.Context) {
	enrollment := &domain.Enrollment{}
	err := c.BindJSON(enrollment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.EnrollmentUsecase.Create(enrollment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, enrollment)
}

func (h *EnrollmentHandler) Update(c *gin.Context) {
	id := c.Param("id")

	// Check if ID is empty
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	enrollment := &domain.Enrollment{}
	err := c.BindJSON(enrollment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	enrollment.ID = idInt

	err = h.EnrollmentUsecase.Update(enrollment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, enrollment)
}

func (h *EnrollmentHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.EnrollmentUsecase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment deleted successfully"})
}

func (h *EnrollmentHandler) GetByStudentID(c *gin.Context) {
	studentID := c.Param("studentID")
	enrollments, err := h.EnrollmentUsecase.GetByStudentID(studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(enrollments) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No enrollments found for the given student ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollments fetched successfully", "data": enrollments})
}

func (h *EnrollmentHandler) GetByCourseID(c *gin.Context) {
	courseID := c.Param("courseID")
	enrollments, err := h.EnrollmentUsecase.GetByCourseID(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(enrollments) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No enrollments found for the given course ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollments fetched successfully", "data": enrollments})
}
