package blogComponent

import (
	m "quilink/internal/models"
)

type IBlogComponent interface {
	GetBlog(id m.IdRequest) (m.Blog, error)
	GetTotalCount(profileId m.IdRequest) (int64, error)
	GetPostBySlug(slug string) (m.Blog, error)
	GetBlogs(profileId m.IdRequest, page, size int) ([]m.Blog, error)
	GetPublishedBlogs(profileId m.IdRequest, page, size int) ([]m.Blog, error)
	CreateBlog(dto m.BlogDto) (m.Blog, error)
	UpdateBlog(id m.IdRequest, dto m.BlogDto) (m.Blog, error)
	TogglePublishStatus(id m.IdRequest) error
	DeleteBlog(id m.IdRequest) error
}
