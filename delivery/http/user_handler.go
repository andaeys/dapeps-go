package http

import (
	"dapeps-go/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService user.UserService
}

func NewUserHandler(service user.UserService) *UserHandler {
	return &UserHandler{UserService: service}
}

func (h *UserHandler) SetupUserRoutes(router *gin.Engine) {
	router.GET("api/v1/users", h.GetUsers)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.UserService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
