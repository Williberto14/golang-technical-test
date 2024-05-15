package http

import (
	"golang-technical-test/internal/domain"
	"golang-technical-test/internal/usecase"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type GradeHandler struct {
	GradeUsecase usecase.IGradeUsecase
	path         string
}

var (
	gradeHandlerInstance *GradeHandler
	gradeHandlerOnce     sync.Once
)

func NewGradeHandler(gradeUsecase usecase.IGradeUsecase, router *gin.Engine) *GradeHandler {
	gradeHandlerOnce.Do(func() {
		gradeHandlerInstance = &GradeHandler{
			GradeUsecase: gradeUsecase,
			path:         "/grades",
		}
		gradeHandlerInstance.setupRoutes(router)
	})
	return gradeHandlerInstance
}

func (h *GradeHandler) setupRoutes(router *gin.Engine) {
	router.GET(h.path, h.GetAll)
	router.GET(h.path+"/:id", h.GetByID)
	router.POST(h.path+"/create", h.Create)
	router.PUT(h.path+"/update/:id", h.Update)
	router.DELETE(h.path+"/delete/:id", h.Delete)
	router.GET(h.path+"/student/:studentID", h.GetByStudentID)
	router.GET(h.path+"/course/:courseID", h.GetByCourseID)
	router.GET(h.path+"/professor/:professorID", h.GetByProfessorID)
}

func (h *GradeHandler) GetAll(c *gin.Context) {
	grades, err := h.GradeUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(grades) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No grades found"})
		return
	}
	c.JSON(http.StatusOK, grades)
}

func (h *GradeHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	grade, err := h.GradeUsecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, grade)
}

func (h *GradeHandler) Create(c *gin.Context) {
	var grade domain.Grade
	err := c.BindJSON(&grade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.GradeUsecase.Create(&grade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, grade)
}

func (h *GradeHandler) Update(c *gin.Context) {
	id := c.Param("id")

	// Check if ID is empty
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	var grade domain.Grade
	err := c.BindJSON(&grade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grade.ID = idInt

	err = h.GradeUsecase.Update(&grade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, grade)
}

func (h *GradeHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.GradeUsecase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Grade deleted successfully"})
}

func (h *GradeHandler) GetByStudentID(c *gin.Context) {
	studentID := c.Param("studentID")
	grades, err := h.GradeUsecase.GetByStudentID(studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(grades) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No grades found for this student"})
		return
	}

	c.JSON(http.StatusOK, grades)
}

func (h *GradeHandler) GetByCourseID(c *gin.Context) {
	courseID := c.Param("courseID")
	grades, err := h.GradeUsecase.GetByCourseID(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(grades) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No grades found for this course"})
		return
	}

	c.JSON(http.StatusOK, grades)
}

func (h *GradeHandler) GetByProfessorID(c *gin.Context) {
	professorID := c.Param("professorID")
	grades, err := h.GradeUsecase.GetByProfessorID(professorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(grades) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No grades found for this professor"})
		return
	}

	c.JSON(http.StatusOK, grades)
}

//obtener notas
