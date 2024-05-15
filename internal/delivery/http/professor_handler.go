package http

import (
	"golang-technical-test/internal/domain"
	"golang-technical-test/internal/usecase"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type ProfessorHandler struct {
	ProfessorUsecase usecase.IProfessorUsecase
	path             string
}

var (
	professorHandlerInstance *ProfessorHandler
	professorHandlerOnce     sync.Once
)

func NewProfessorHandler(professorUsecase usecase.IProfessorUsecase, router *gin.Engine) *ProfessorHandler {
	professorHandlerOnce.Do(func() {
		professorHandlerInstance = &ProfessorHandler{
			ProfessorUsecase: professorUsecase,
			path:             "/professor",
		}
		professorHandlerInstance.setupRoutes(router)
	})
	return professorHandlerInstance
}

func (h *ProfessorHandler) setupRoutes(router *gin.Engine) {
	router.GET(h.path, h.GetAll)
	router.GET(h.path+"/:id", h.GetByID)
	router.POST(h.path+"/create", h.Create)
	router.PUT(h.path+"/update/:id", h.Update)
	router.DELETE(h.path+"/delete/:id", h.Delete)
}

func (h *ProfessorHandler) GetAll(c *gin.Context) {
	professors, err := h.ProfessorUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(professors) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No professors found"})
		return
	}

	c.JSON(http.StatusOK, professors)
}

func (h *ProfessorHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return

	}

	professor, err := h.ProfessorUsecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, professor)
}

func (h *ProfessorHandler) Create(c *gin.Context) {
	var professor domain.Professor
	if err := c.ShouldBindJSON(&professor); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.ProfessorUsecase.Create(&professor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, professor)
}

func (h *ProfessorHandler) Update(c *gin.Context) {
	id := c.Param("id")

	// Check if ID is empty
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	var professor domain.Professor
	if err := c.ShouldBindJSON(&professor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	professor.ID = idInt

	if err := h.ProfessorUsecase.Update(&professor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, professor)
}

func (h *ProfessorHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := h.ProfessorUsecase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Professor deleted successfully"})
}
