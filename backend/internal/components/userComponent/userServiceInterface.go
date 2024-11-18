package userComponent

import (
	m "quilink/internal/models"

	"github.com/google/uuid"
)

type IUserService interface {
	Register(dto m.UserDto) (uuid.UUID, error)
	Login(email, password string) (uuid.UUID, error)
	DeleteAccount(id uuid.UUID) (bool, error)
}
