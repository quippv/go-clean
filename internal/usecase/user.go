package usecase

import (
	"github.com/google/uuid"
	"github.com/quippv/go-clean/internal/entity"
)

// defines the interface for user repository contract
type UserRepository interface {
	GetUserById(id uuid.UUID) (*entity.User, error)
	CreateUser(user *entity.User) (uuid.UUID, error)
}

type UserUseCase struct {
	userRepo        UserRepository
	passwordUseCase *PasswordUseCase
}

// creates a new instance of UserUseCase
func NewUserUseCase(userRepo UserRepository, passwordUseCase *PasswordUseCase) *UserUseCase {
	return &UserUseCase{userRepo, passwordUseCase}
}

func (u *UserUseCase) GetUser(id uuid.UUID) (*entity.User, error) {
	return u.userRepo.GetUserById(id)
}

func (u *UserUseCase) RegisterUser(user *entity.User) (uuid.UUID, error) {
	// Hash the password before storing it
	hashedPassword, err := u.passwordUseCase.HashPassword(user.Password)
	if err != nil {
		return uuid.Nil, err
	}

	// Set the hashed password to the user object
	user.Password = hashedPassword

	// Create the user in the repository
	return u.userRepo.CreateUser(user)
}
