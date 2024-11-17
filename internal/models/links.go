package models

import "github.com/google/uuid"

type Link struct {
	Super
	Title  string
	URL    string
	Order  int
	Active bool

	// relationship
	UserId uuid.UUID
}
