package userComponent

import (
	"fmt"
	"log"
	m "quilink/internal/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db gorm.DB
}

func NewUserService(db gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) Register(dto m.UserDto) (uuid.UUID, error) {
	hash, err := s.hashPassword(dto.Password)
	if err != nil {
		return uuid.UUID{}, err
	}

	user := m.User{
		Username:      dto.Username,
		Password_hash: hash,
		Email:         dto.Email,
	}

	result := s.db.Create(&user)

	if result.Error != nil {
		log.Printf("[UserService.Register] error creating user: %v", result.Error)
		return uuid.UUID{}, fmt.Errorf("failed to create user %w", result.Error)
	}

	return user.ID, nil
}

func (s *UserService) Login(email, password string) (uuid.UUID, error) {
	var user m.User

	if err := s.db.First(&user, "email = ?", email).Error; err != nil {
		log.Printf("[UserService.Login] error finding user with email %s: %v", email, err)
		return uuid.UUID{}, fmt.Errorf("failed to find user %w", err)
	}

	hash := user.Password_hash

	if check := s.checkHash(password, hash); !check {
		log.Printf("[UserService.Login] password does not match")
		return uuid.UUID{}, fmt.Errorf("failed to match password")
	}

	return user.ID, nil
}

func (s *UserService) DeleteAccount(id uuid.UUID) (bool, error) {
	if err := s.db.Delete(&m.User{}, id).Error; err != nil {
		log.Printf("[UserService.DeleteAccount] failed to delete user with id %s: %v", id, err)
		return false, fmt.Errorf("failed to delete user %s: %v", id, err)
	}

	return true, nil
}

func (s *UserService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil

}

func (s *UserService) checkHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
