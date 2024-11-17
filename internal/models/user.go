package models

type User struct {
	Super
	Username      string
	Password_hash string
	Email         string

	// relationships
	Profile Profile `gorm:"foreignKey:UserId"`
	Blogs   []Blog  `gorm:"foreignKey:UserId"`
	Link    []Link  `gorm:"foreignKey:UserId"`
}
