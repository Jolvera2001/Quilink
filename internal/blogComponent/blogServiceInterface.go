package blogComponent

import (
	m "quilink/internal/models"

	"github.com/google/uuid"
)

type IBlogComponent interface {
	GetBlog(id uuid.UUID) (m.Blog, error)
	GetBlogs(userId uuid.UUID, page, size int) ([]m.Blog, error)
}
