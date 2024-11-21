package handlers

import (
	"log"
	"net/http"
	c "quilink/internal/components/userComponent"
	m "quilink/internal/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service c.IUserService
}

func NewUserHandler(service c.IUserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func GroupUserHandlers(r *gin.Engine, h *UserHandler) {
	v1 := r.Group("api/v1")
	{
		v1.POST("account/register", h.Register)
		v1.POST("account/login", h.Login)
		v1.DELETE("account/delete", h.DeleteAccount)
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var userDto m.UserDto

	if err := c.ShouldBindJSON(&userDto); err != nil {
		log.Printf("[UserHandler.Reigster] invalid data recieved: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	userId, err := h.service.Register(userDto)
	if err != nil {
		log.Printf("[UserHandler.Reigster] error registering new user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"account_id": userId})
}

func (h *UserHandler) Login(c *gin.Context) {
	var userLoginDto m.UserLoginDto

	if err := c.ShouldBindJSON(&userLoginDto); err != nil {
		log.Printf("[UserHandler.Login] invalid data recieved: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	userId, err := h.service.Login(userLoginDto)
	if err != nil {
		log.Printf("[UserHandler.Login] error logging in user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"account_id": userId})
}

func (h *UserHandler) DeleteAccount(c *gin.Context) {
	var id m.IdRequest

	if err := c.ShouldBindJSON(&id); err != nil {
		log.Printf("[UserHandler.DeleteAccount] invalid data recieved: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	deleted, err := h.service.DeleteAccount(id)
	if err != nil {
		log.Printf("[UserHandler.DeleteAccount] error deleting user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": deleted})

}
