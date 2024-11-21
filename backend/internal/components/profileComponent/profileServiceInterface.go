package profilecomponent

import (
	m "quilink/internal/models"

	"github.com/google/uuid"
)

type IProfileService interface {
	GetProfile(id uuid.UUID) (m.Profile, []m.Link, error)
	GetProfiles(userId uuid.UUID) ([]m.Profile, error)
	GetByDomain(customDomain string) (m.Profile, []m.Link, error)
	CreateProfile(dto m.ProfileDto) (m.Profile, error)
	UpdateProfile(id uuid.UUID, dto m.ProfileDto) (m.Profile, error) 
	DeleteProfile(id uuid.UUID) error
}
