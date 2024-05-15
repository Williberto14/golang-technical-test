package http

import (
	"golang-technical-test/internal/domain"
	"golang-technical-test/internal/usecase"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type CoursesHandler struct {
	CoursesUsecase usecase.ICourseUsecase
	path           string
}

var (
	coursesHandlerInstance *CoursesHandler
	coursesHandlerOnce     sync.Once
)

func NewCourseHandler(coursesUsecase usecase.ICourseUsecase, router *gin.Engine) *CoursesHandler {
	coursesHandlerOnce.Do(func() {
		coursesHandlerInstance = &CoursesHandler{
			CoursesUsecase: coursesUsecase,
			path:           "/courses",
		}
		coursesHandlerInstance.setupRoutes(router)
	})
	return coursesHandlerInstance
}

func (h *CoursesHandler) setupRoutes(router *gin.Engine) {
	router.GET(h.path, h.GetAll)
	router.GET(h.path+"/:id", h.GetByID)
	router.POST(h.path+"/create", h.Create)
	router.PUT(h.path+"/update/:id", h.Update)
	router.DELETE(h.path+"/delete/:id", h.Delete)
}

func (h *CoursesHandler) GetAll(c *gin.Context) {
	courses, err := h.CoursesUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func (h *CoursesHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	courses, err := h.CoursesUsecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func (h *CoursesHandler) Create(c *gin.Context) {
	var courses domain.Course
	if err := c.ShouldBindJSON(&courses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.CoursesUsecase.Create(&courses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, courses)
}

func (h *CoursesHandler) Update(c *gin.Context) {
	id := c.Param("id")

	// Check if ID is empty
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	var courses domain.Course
	if err := c.ShouldBindJSON(&courses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	courses.ID = idInt

	if err := h.CoursesUsecase.Update(&courses); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func (h *CoursesHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.CoursesUsecase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}
