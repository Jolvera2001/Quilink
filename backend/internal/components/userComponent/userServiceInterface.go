package userComponent

import (
	m "quilink/internal/models"

	"github.com/google/uuid"
)

type IUserService interface {
	Register(dto m.UserDto) (uuid.UUID, error)
	Login(loginDto m.UserLoginDto) (uuid.UUID, error)
	DeleteAccount(id m.IdRequest) (bool, error)
}
