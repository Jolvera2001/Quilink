package models

import "github.com/google/uuid"

type Blog struct {
	Super
	Title     string `gorm:"not null"`
	Content   string `gorm:"type:text"`
	Slug      string
	Published bool `gorm:"default:false"`

	// relationship
	UserId uuid.UUID `gorm:"type:uuid;not null"`
}

type BlogDto struct {
	Title     string
	Content   string
	Slug      string
	Published bool
	UserId    uuid.UUID
}
