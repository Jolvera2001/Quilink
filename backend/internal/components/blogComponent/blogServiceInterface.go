package blogComponent

import (
	m "quilink/internal/models"

	"github.com/google/uuid"
)

type IBlogComponent interface {
	GetBlog(id uuid.UUID) (m.Blog, error)
	GetTotalCount(userId uuid.UUID) (int64, error)
	GetPostBySlug(slug string) (m.Blog, error)
	GetBlogs(userId uuid.UUID, page, size int) ([]m.Blog, error)
	GetPublishedBlogs(userId uuid.UUID, page, size int) ([]m.Blog, error)
	CreateBlog(dto m.BlogDto) (m.Blog, error)
	UpdateBlog(blogId uuid.UUID, dto m.BlogDto) (m.Blog, error)
	TogglePublishStatus(id uuid.UUID) error
	DeleteBlog(id uuid.UUID) error
}
