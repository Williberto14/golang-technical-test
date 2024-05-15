package http

import (
	"golang-technical-test/internal/domain"
	"golang-technical-test/internal/usecase"
	"net/http"
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
	router.GET(h.path, h.GetAll)
	router.GET(h.path+"/:id", h.GetByID)
	router.POST(h.path, h.Create)
	router.PUT(h.path, h.Update)
	router.DELETE(h.path+"/:id", h.Delete)
	router.GET(h.path+"/student/:studentID", h.GetByStudentID)
	router.GET(h.path+"/course/:courseID", h.GetByCourseID)
}

func (h *EnrollmentHandler) GetAll(c *gin.Context) {
	enrollments, err := h.EnrollmentUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	enrollment := &domain.Enrollment{}
	err := c.BindJSON(enrollment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	c.JSON(http.StatusOK, enrollments)
}

func (h *EnrollmentHandler) GetByCourseID(c *gin.Context) {
	courseID := c.Param("courseID")
	enrollments, err := h.EnrollmentUsecase.GetByCourseID(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, enrollments)
}
