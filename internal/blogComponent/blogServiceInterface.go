package blogComponent

import "github.com/google/uuid"

type IBlogComponent interface {
	GetBlog(id uuid.UUID)
}
