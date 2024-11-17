package models

import (
	"time"

	"github.com/google/uuid"
)

type Super struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}
