package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Super struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}

func (s *Super) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}

	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}

	if s.UpdatedAt.IsZero() {
		s.UpdatedAt = s.CreatedAt
	}

	return nil
}

func (s *Super) BeforeUpdate(tx *gorm.DB) error {
	s.UpdatedAt = time.Now()
	return nil
}
