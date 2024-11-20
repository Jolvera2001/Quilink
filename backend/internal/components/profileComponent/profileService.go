package profilecomponent

import "gorm.io/gorm"

type ProfileService struct {
	db *gorm.DB
}

func NewProfileService(db *gorm.DB) *ProfileService {
	return &ProfileService{
		db: db,
	}
}