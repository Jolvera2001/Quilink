package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	Placeholder int
}

func NewBlogHandler() *BlogHandler {
	return &BlogHandler{
		Placeholder: 1,
	}
}

func (h *BlogHandler) GetInt(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"int": h.Placeholder})
}

func GroupBlogHandlers(r *gin.Engine, h *BlogHandler) {
	v1 := r.Group("api/v1")
	{
		v1.GET("getint", h.GetInt)
	}
}
