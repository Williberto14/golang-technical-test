package http

import (
	"golang-technical-test/middlewares"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	path string
}

var (
	loginHandlerInstance *LoginHandler
	loginHandlerOnce     sync.Once
)

func NewLoginHandler(router *gin.Engine) *LoginHandler {
	loginHandlerOnce.Do(func() {
		loginHandlerInstance = &LoginHandler{
			path: "/login",
		}
		loginHandlerInstance.setupRoutes(router)
	})
	return loginHandlerInstance
}

func (h *LoginHandler) setupRoutes(router *gin.Engine) {
	router.POST(h.path, h.Login)
}

func (h *LoginHandler) Login(c *gin.Context) {
	var login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if login.Username == "test" && login.Password == "test" {
		token, err := middlewares.CreateToken(login.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
