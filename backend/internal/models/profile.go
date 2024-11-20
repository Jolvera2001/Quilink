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

type ProfileDto struct {
	Display_name  string    `json:"display_name"`
	Bio           string    `json:"bio"`
	Avatar_url    string    `json:"avatar_url"`
	Theme         string    `json:"theme"`
	Custom_domain string    `json:"custom_domain"`
	UserId        uuid.UUID `json:"user_id"`
}
