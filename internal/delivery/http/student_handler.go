package http

import (
	"golang-technical-test/internal/domain"
	"golang-technical-test/internal/usecase"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	StudentUsecase usecase.IStudentUsecase
	path           string
}

var (
	studentHandlerInstance *StudentHandler
	studentHandlerOnce     sync.Once
)

func NewStudentHandler(studentUsecase usecase.IStudentUsecase, router *gin.Engine) *StudentHandler {
	studentHandlerOnce.Do(func() {
		studentHandlerInstance = &StudentHandler{
			StudentUsecase: studentUsecase,
			path:           "/students",
		}
		studentHandlerInstance.setupRoutes(router)
	})
	return studentHandlerInstance
}

func (h *StudentHandler) setupRoutes(router *gin.Engine) {
	router.GET(h.path, h.GetAll)
	router.GET(h.path+"/:id", h.GetByID)
	router.POST(h.path+"/create", h.Create)
	router.PUT(h.path+"/update/:id", h.Update)
	router.DELETE(h.path+"/delete/:id", h.Delete)
}

func (h *StudentHandler) GetAll(c *gin.Context) {
	students, err := h.StudentUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	student, err := h.StudentUsecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) Create(c *gin.Context) {
	var student domain.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.StudentUsecase.Create(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, student)
}

func (h *StudentHandler) Update(c *gin.Context) {
	id := c.Param("id")

	// Check if ID is empty
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	var student domain.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student.ID = idInt

	if err := h.StudentUsecase.Update(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.StudentUsecase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
