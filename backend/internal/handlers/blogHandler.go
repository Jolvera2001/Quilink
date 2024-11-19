package handlers

import (
	c "quilink/internal/components/blogComponent"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	Service *c.BlogService
}

func NewBlogHandler(service *c.BlogService) *BlogHandler {
	return &BlogHandler{
		Service: service,
	}
}

func GroupBlogHandlers(r *gin.Engine, h *BlogHandler) {
	v1 := r.Group("api/v1")
	{
		v1.GET("blog/all/:profileId", h.GetBlogs)
		v1.GET("blog/all/published/:profileId", h.GetPublishedBlogs)
		v1.GET("blog/total/:profileId", h.GetTotalCount)

		v1.POST("blog", h.CreateBlog)
		v1.PUT("blog/:id", h.UpdateBlog)
		v1.PATCH("blog/:id/publish", h.UpdateBlog)
		v1.DELETE("blog/:id", h.DeleteBlog)

		v1.GET("blog/by-slug/:slug", h.GetPostBySlug)
		v1.GET("blog/:id", h.GetBlog)
	}
}

func (h *BlogHandler) GetBlog(c *gin.Context) {

}

func (h *BlogHandler) GetTotalCount(c *gin.Context) {

}

func (h *BlogHandler) GetPostBySlug(c *gin.Context) {

}

func (h *BlogHandler) GetBlogs(c *gin.Context) {

}

func (h *BlogHandler) GetPublishedBlogs(c *gin.Context) {

}

func (h *BlogHandler) CreateBlog(c *gin.Context) {

}

func (h *BlogHandler) UpdateBlog(c *gin.Context) {

}

func (h *BlogHandler) TogglePublishStatus(c *gin.Context) {

}

func (h *BlogHandler) DeleteBlog(c *gin.Context) {

}
