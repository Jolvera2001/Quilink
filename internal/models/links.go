package models

import "github.com/google/uuid"

type Link struct {
	Super
	Title  string `gorm:"not null"`
	URL    string `gorm:"not null"`
	Order  int
	Active bool `gorm:"default:true"`

	// relationship
	UserId uuid.UUID `gorm:"type:uuid;not null"`
}
