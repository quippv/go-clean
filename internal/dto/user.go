package dto

import (
	"github.com/google/uuid"
	"github.com/quippv/go-clean/internal/entity"
)

// CreateUserDTO represents the payload for creating a new user
type CreateUserDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// UserResponseDTO represents the data returned for a user
type UserResponseDTO struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

// ToEntity converts CreateUserDTO to entity.User
func (dto *CreateUserDTO) ToEntity() *entity.User {
	return &entity.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

// ToDTO converts entity.User to UserResponseDTO
func ToDTO(user *entity.User) *UserResponseDTO {
	return &UserResponseDTO{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: "xxx",
	}
}
