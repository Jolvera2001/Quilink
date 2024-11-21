package handlers

import (
	c "quilink/internal/components/profileComponent"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	service c.IProfileService
}

func NewProfileHandler(service c.IProfileService) *ProfileHandler {
	return &ProfileHandler{
		service: service,
	}
}

func GroupProfileHandlers(r *gin.Engine, h ProfileHandler) {
	v1 := r.Group("api/v1") 
	{
		v1.GET("profile/:id", h.GetProfile)
		v1.GET("profile/all/:userId", h.GetProfiles)
		v1.GET("profile/name/:domain", h.GetByDomain)

		v1.POST("profile", h.CreateProfile)
		v1.PUT("profile", h.UpdateProfile)
		v1.DELETE("profile/:id", h.DeleteProfile)
	}
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {

}

func (h *ProfileHandler) GetProfiles(c *gin.Context) {

}

func (h *ProfileHandler) GetByDomain(c *gin.Context) {

}

func (h *ProfileHandler) CreateProfile(c *gin.Context) {

}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {

}

func (h *ProfileHandler) DeleteProfile(c *gin.Context) {

}




