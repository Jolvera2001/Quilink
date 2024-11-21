package handlers

import (
	"log"
	"net/http"
	c "quilink/internal/components/blogComponent"
	m "quilink/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BlogHandler struct {
	service *c.BlogService
}

func NewBlogHandler(service *c.BlogService) *BlogHandler {
	return &BlogHandler{
		service: service,
	}
}

func GroupBlogHandlers(r *gin.Engine, h *BlogHandler) {
	v1 := r.Group("api/v1")
	{
		v1.GET("blog/all/:profileId/:page/:pageSize", h.GetBlogs)
		v1.GET("blog/all/published/:profileId/:page/:pageSize", h.GetPublishedBlogs)
		v1.GET("blog/total/:profileId", h.GetTotalCount)

		v1.POST("blog", h.CreateBlog)
		v1.PUT("blog/:id", h.UpdateBlog)
		v1.PATCH("blog/:id/publish", h.TogglePublishStatus)
		v1.DELETE("blog/:id", h.DeleteBlog)

		v1.GET("blog/by-slug/:slug", h.GetPostBySlug)
		v1.GET("blog/:id", h.GetBlog)
	}
}

func (h *BlogHandler) GetBlog(c *gin.Context) {
	blogId, err := uuid.Parse(c.Param("id"))

	if err != nil {
		log.Printf("[BlogHandler.GetBlog] unable to parse uuid: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	blog, err := h.service.GetBlog(blogId)
	if err != nil {
		log.Printf("[BlogHandler.GetBlog] error getting blog: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blog": blog})
}

func (h *BlogHandler) GetTotalCount(c *gin.Context) {
	profileId, err := uuid.Parse(c.Param("profileId"))

	if err != nil {
		log.Printf("[BlogHandler.GetTotalCount] unable to parse uuid: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	count, err := h.service.GetTotalCount(profileId)
	if err != nil {
		log.Printf("[BlogHandler.GetTotalCount] error getting total blog count: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total_count": count})
}

func (h *BlogHandler) GetPostBySlug(c *gin.Context) {
	slug := c.Param("slug")

	blog, err := h.service.GetPostBySlug(slug)
	if err != nil {
		log.Printf("[BlogHandler.GetPostBySlug] error getting blog from slug: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blog": blog})
}

func (h *BlogHandler) GetBlogs(c *gin.Context) {
	profileId, err := uuid.Parse(c.Param("profileId"))

	if err != nil {
		log.Printf("[BlogHandler.GetBlogs] unable to parse uuid: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	page, err := strconv.Atoi(c.Param("page"))

	if err != nil {
		log.Printf("[BlogHandler.GetBlogs] unable to parse page: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	pageSize, err := strconv.Atoi(c.Param("pageSize"))

	if err != nil {
		log.Printf("[BlogHandler.GetBlogs] unable to parse pageNumber: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	blogs, err := h.service.GetBlogs(profileId, page, pageSize)
	if err != nil {
		log.Printf("[BlogHandler.GetBlogs] error getting blogs: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blogs": blogs})
}

func (h *BlogHandler) GetPublishedBlogs(c *gin.Context) {
	profileId, err := uuid.Parse(c.Param("profileId"))

	if err != nil {
		log.Printf("[BlogHandler.GetPublishedBlogs] unable to parse uuid: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	page, err := strconv.Atoi(c.Param("page"))

	if err != nil {
		log.Printf("[BlogHandler.GetPublishedBlogs] unable to parse page: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	pageSize, err := strconv.Atoi(c.Param("pageSize"))

	if err != nil {
		log.Printf("[BlogHandler.GetPublishedBlogs] unable to parse pageNumber: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	blogs, err := h.service.GetPublishedBlogs(profileId, page, pageSize)
	if err != nil {
		log.Printf("[BlogHandler.GetPublishedBlogs] error getting published blogs: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blogs": blogs})
}

func (h *BlogHandler) CreateBlog(c *gin.Context) {
	var blog m.BlogDto

	if err := c.ShouldBindJSON(&blog); err != nil {
		log.Printf("[BlogHandler.CreateBlog] invalid data recieved: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	createdBlog, err := h.service.CreateBlog(blog)
	if err != nil {
		log.Printf("[BlogHandler.CreateBlog] error creating blog: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"blog": createdBlog})
}

func (h *BlogHandler) UpdateBlog(c *gin.Context) {
	var blogUpdate m.BlogDto
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		log.Printf("[BlogHandler.UpdateBlog] unable to parse uuid: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := c.ShouldBindJSON(&blogUpdate); err != nil {
		log.Printf("[BlogHandler.UpdateBlog] invalid data recieved: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	blog, err := h.service.UpdateBlog(id, blogUpdate)
	if err != nil {
		log.Printf("[BlogHandler.UpdateBlog] error updating blog: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blog": blog})
}

func (h *BlogHandler) TogglePublishStatus(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		log.Printf("[BlogHandler.TogglePublishStatus] unable to parse uuid: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	check, err := h.service.TogglePublishStatus(id)
	if err != nil {
		log.Printf("[BlogHandler.TogglePublishStatus] error toggling blog being published: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"published_status": check})
}

func (h *BlogHandler) DeleteBlog(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		log.Printf("[BlogHandler.DeleteBlog] unable to parse uuid: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = h.service.DeleteBlog(id)
	if err != nil {
		log.Printf("[BlogHandler.DeleteBlog] error deleting blog: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, nil)

}
