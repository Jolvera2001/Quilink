package models

import "github.com/google/uuid"

type Blog struct {
	Super
	Title     string `gorm:"not null"`
	Content   string `gorm:"type:text"`
	Slug      string
	Published bool `gorm:"default:false"`

	// relationship
	ProfileId uuid.UUID `gorm:"type:uuid;not null"`
}

type BlogDto struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Slug      string    `json:"slug"`
	Published bool      `json:"published"`
	ProfileId    uuid.UUID `json:"user_id"`
}
