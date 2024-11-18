package models

type User struct {
	Super
	Username      string `gorm:"uniqueIndex;not null"`
	Password_hash string
	Email         string `gorm:"uniqueIndex;not null"`

	// relationships
	Profile Profile `gorm:"foreignKey:UserId"`
	Blogs   []Blog  `gorm:"foreignKey:UserId"`
	Link    []Link  `gorm:"foreignKey:UserId"`
}

type UserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserLoginDto struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
