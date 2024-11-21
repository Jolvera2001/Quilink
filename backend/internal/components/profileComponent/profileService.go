package profilecomponent

import (
	"fmt"
	"log"
	m "quilink/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProfileService struct {
	db *gorm.DB
}

func NewProfileService(db *gorm.DB) *ProfileService {
	return &ProfileService{
		db: db,
	}
}

func (s *ProfileService) GetProfile(id uuid.UUID) (m.Profile, error) {
	var profile m.Profile

	if err := s.db.First(&profile, "id = ?", id).Error; err != nil {
		log.Printf("[ProfileSerivce.GetProfile][profileId=%s] error finding profile with id %s: %v", id, id, err)
		if err == gorm.ErrRecordNotFound {
			return m.Profile{}, fmt.Errorf("profile not found with id %s", id)
		}
		return m.Profile{}, fmt.Errorf("error fetching profile from database: %w", err)
	}

	return profile, nil
}

func (s *ProfileService) GetProfiles(userId uuid.UUID) ([]m.Profile, error) {
	var profiles []m.Profile

	result := s.db.
		Where("user_id = ?", userId).
		Order("created_at DESC").
		Find(&profiles)

	if result.Error != nil {
		log.Printf("[ProfileSerivce.GetProfiles][userId=%s] error finding profile with id %s: %v", userId, userId, result.Error)
		return []m.Profile{}, nil
	}

	return profiles, nil
}

func (s *ProfileService) GetByDomain(customDomain string) (m.Profile, error) {
	var profile m.Profile

	if err := s.db.First(&profile, "custom_domain = ?", customDomain).Error; err != nil {
		log.Printf("[ProfileSerivce.GetByDomain][domain=%s] error finding profile with domain %s: %v", customDomain, customDomain, err)
		if err == gorm.ErrRecordNotFound {
			return m.Profile{}, fmt.Errorf("profile not found with domain %s", customDomain)
		}
		return m.Profile{}, fmt.Errorf("error fetching profile from database: %w", err)
	}

	return profile, nil
}

func (s *ProfileService) CreateProfile(dto m.ProfileDto) (m.Profile, error) {
	profile := m.Profile{
		Display_name: dto.Display_name,
		Bio: dto.Bio,
		Avatar_url: dto.Avatar_url,
		Theme: dto.Theme,
		Custom_domain: dto.Custom_domain,
		UserId: dto.UserId,
	}

	result := s.db.Create(&profile)

	if result.Error != nil {
		log.Printf("[ProfileSerivce.CreateProfile] error creating profile: %v", result.Error)
		return m.Profile{}, result.Error
	}

	return profile, nil
}

func (s *ProfileService) UpdateProfile(id uuid.UUID, dto m.ProfileDto) (m.Profile, error) {
	var profileToUpdate m.Profile

	if err := s.db.First(&profileToUpdate, "id = ?", id).Error; err != nil {
		log.Printf("[ProfileSerivce.UpdateProfile][id=%s] error finding profile with id %s: %v", id, id, err)
		if err == gorm.ErrRecordNotFound {
			return m.Profile{}, fmt.Errorf("profile not found with id %s", id)
		}
		return m.Profile{}, fmt.Errorf("error fetching profile from database: %w", err)
	}

	profileToUpdate.Display_name = dto.Display_name
	profileToUpdate.Bio = dto.Bio
	profileToUpdate.Avatar_url = dto.Avatar_url
	profileToUpdate.Theme = dto.Theme
	profileToUpdate.Custom_domain = dto.Custom_domain

	if err := s.db.Save(&profileToUpdate).Error; err != nil {
		log.Printf("[ProfileService.UpdateProfile] error updating profile: id=%s error=%v", id, err)
		return m.Profile{}, fmt.Errorf("failed to update profile: %w", err)
	}

	return profileToUpdate, nil
}

func (s *ProfileService) DeleteProfile(id uuid.UUID) error {
	if err := s.db.Delete(&m.Profile{}, id).Error; err != nil {
		log.Printf("[ProfileService.DeleteProfile] error deleting profile: id=%s error=%v", id, err)
		return fmt.Errorf("failed to delete profile: %w", err)
	}

	return nil
}
