package http

import (
	"golang-technical-test/internal/domain"
	"golang-technical-test/internal/usecase"
	"net/http"
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
	router.POST(h.path, h.Create)
	router.PUT(h.path, h.Update)
	router.DELETE(h.path+"/:id", h.Delete)
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
	var grade domain.Grade
	err := c.BindJSON(&grade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
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
	c.JSON(http.StatusNoContent, nil)
}

func (h *GradeHandler) GetByStudentID(c *gin.Context) {
	studentID := c.Param("studentID")
	grades, err := h.GradeUsecase.GetByStudentID(studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	c.JSON(http.StatusOK, grades)
}

func (h *GradeHandler) GetByProfessorID(c *gin.Context) {
	professorID := c.Param("professorID")
	grades, err := h.GradeUsecase.GetByProfessorID(professorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, grades)
}

//obtener notas
