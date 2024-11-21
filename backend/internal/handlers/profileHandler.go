package handlers

import (
	"log"
	"net/http"
	c "quilink/internal/components/profileComponent"
	m "quilink/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProfileHandler struct {
	service c.IProfileService
}

func NewProfileHandler(service c.IProfileService) *ProfileHandler {
	return &ProfileHandler{
		service: service,
	}
}

func GroupProfileHandlers(r *gin.Engine, h *ProfileHandler) {
	v1 := r.Group("api/v1")
	{
		v1.GET("profile/:id", h.GetProfile)
		v1.GET("profile/all/:userId", h.GetProfiles)
		v1.GET("profile/name/:domain", h.GetByDomain)

		v1.POST("profile", h.CreateProfile)
		v1.PUT("profile/:id", h.UpdateProfile)
		v1.DELETE("profile/:id", h.DeleteProfile)
	}
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Printf("[ProfileHandler.GetProfile] unable to parse uuid: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	profile, err := h.service.GetProfile(id)
	if err != nil {
		log.Printf("[ProfileHandler.GetProfile] error getting profile: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": profile})
}

func (h *ProfileHandler) GetProfiles(c *gin.Context) {
	userId, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		log.Printf("[ProfileHandler.GetProfiles] unable to parse uuid: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	profiles, err := h.service.GetProfiles(userId)
	if err != nil {
		log.Printf("[ProfileHandler.GetProfiles] error getting profiles: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profiles": profiles})
}

func (h *ProfileHandler) GetByDomain(c *gin.Context) {
	domain := c.Param("domain")

	profile, err := h.service.GetByDomain(domain)
	if err != nil {
		log.Printf("[ProfileHandler.GetByDomain] error getting profile: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": profile})
}

func (h *ProfileHandler) CreateProfile(c *gin.Context) {
	var profileDto m.ProfileDto

	if err := c.ShouldBindJSON(&profileDto); err != nil {
		log.Printf("[ProfileHandler.CreateProfile] received invalid data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	profile, err := h.service.CreateProfile(profileDto)
	if err != nil {
		log.Printf("[ProfileHandler.CreateProfile] error creating profile: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": profile})
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	var profileDto m.ProfileDto
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Printf("[ProfileHandler.GetProfiles] unable to parse uuid: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := c.ShouldBindJSON(&profileDto); err != nil {
		log.Printf("[ProfileHandler.UpdateProfile] received invalid data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	profile, err := h.service.UpdateProfile(id, profileDto)
	if err != nil {
		log.Printf("[ProfileHandler.UpdateProfile] error updating profile: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": profile})
}

func (h *ProfileHandler) DeleteProfile(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Printf("[ProfileHandler.DeleteProfile] unable to parse uuid: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := h.service.DeleteProfile(id); err != nil {
		log.Printf("[ProfileHandler.DeleteProfile] error deleting profile: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": "success"})
}
