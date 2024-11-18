package models

import "github.com/google/uuid"

type Profile struct {
	Super
	Display_name  string
	Bio           string
	Avatar_url    string
	Theme         string
	Custom_domain string

	// relationship
	UserId uuid.UUID `gorm:"type:uuid;not null"`
	Blogs  []Blog    `gorm:"foreignKey:ProfileId"`
	Link   []Link    `gorm:"foreignKey:ProfileId"`
}
