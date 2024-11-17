package models

import "github.com/google/uuid"

type Blog struct {
	Super
	Title     string
	Content   string
	Slug      string
	Published bool

	// relationship
	UserId uuid.UUID
}
