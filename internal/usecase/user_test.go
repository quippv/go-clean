package usecase_test

import (
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/quippv/go-clean/internal/entity"
	"github.com/quippv/go-clean/internal/usecase"
	"github.com/quippv/go-clean/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	saltLength = 16
	keyLength  = 32
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUserById(id uuid.UUID) (*entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) CreateUser(user *entity.User) (uuid.UUID, error) {
	args := m.Called(user)
	return args.Get(0).(uuid.UUID), args.Error(1)
}

func TestGetUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	passwordUseCase := usecase.NewPasswordUseCase()
	useCase := usecase.NewUserUseCase(mockRepo, passwordUseCase)
	id, createdAt, updatedAt := utils.GenerateIDAndUnixMillis()
	testUser := &entity.User{ID: id, Name: faker.Name(), Email: faker.Email(), Password: faker.Password(), CreatedAt: createdAt, UpdatedAt: updatedAt}

	mockRepo.On("GetUserById", id).Return(testUser, nil)

	user, err := useCase.GetUser(id)
	assert.NoError(t, err)
	assert.Equal(t, testUser, user)
	mockRepo.AssertExpectations(t)
}

func TestRegisterUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	passwordUseCase := usecase.NewPasswordUseCase()
	useCase := usecase.NewUserUseCase(mockRepo, passwordUseCase)
	testId := uuid.New()
	testUser := &entity.User{ID: testId, Name: faker.Name(), Email: faker.Email(), Password: faker.Password()}

	hashedPassword, err := passwordUseCase.HashPassword(testUser.Password)
	assert.NoError(t, err)
	testUser.Password = hashedPassword

	mockRepo.On("CreateUser", testUser).Return(testId, nil)

	id, err := useCase.RegisterUser(testUser)
	assert.NoError(t, err)
	assert.Equal(t, testId, id)
	mockRepo.AssertExpectations(t)
}
