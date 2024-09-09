package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"strings"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/quippv/go-clean/internal/dto"
	"github.com/quippv/go-clean/internal/entity"
	httpHandler "github.com/quippv/go-clean/internal/handler/http"
	"github.com/quippv/go-clean/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserUseCase is a mock implementation of UserUseCase
type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) GetUser(id uuid.UUID) (*entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserUseCase) RegisterUser(user *entity.User) (uuid.UUID, error) {
	args := m.Called(user)
	return args.Get(0).(uuid.UUID), args.Error(1)
}

func TestGetUser(t *testing.T) {
	e := echo.New()

	// Create a new instance of the mock use case
	mockUseCase := new(MockUserUseCase)

	userID, _, _ := utils.GenerateIDAndUnixMillis()
	expectedUser := &entity.User{
		ID:       userID,
		Name:     faker.Name(),
		Email:    faker.Email(),
		Password: faker.Password(),
	}

	handler := &httpHandler.UserHandler{UserUseCase: mockUseCase}
	httpHandler.NewUserHttpHandler(e, mockUseCase)
	mockUseCase.On("GetUser", userID).Return(expectedUser, nil)

	req := httptest.NewRequest(echo.GET, "/users/"+userID.String(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("users/:id")
	c.SetParamNames("id")
	c.SetParamValues(userID.String())

	// Call the handler
	if assert.NoError(t, handler.GetUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var userResponse dto.UserResponseDTO
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &userResponse)) {
			assert.Equal(t, expectedUser.ID, userResponse.ID)
			assert.Equal(t, expectedUser.Name, userResponse.Name)
			assert.Equal(t, expectedUser.Email, userResponse.Email)
			assert.Equal(t, "xxx", userResponse.Password)
		}
	}

	mockUseCase.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	e := echo.New()

	// Create a new instance of the mock use case
	mockUseCase := new(MockUserUseCase)
	userDTO := dto.CreateUserDTO{Name: faker.Name(), Email: faker.Email(), Password: faker.Password()}
	user := &entity.User{Name: userDTO.Name, Email: userDTO.Email, Password: userDTO.Password}
	userID, _, _ := utils.GenerateIDAndUnixMillis()
	mockUseCase.On("RegisterUser", user).Return(userID, nil)

	handler := &httpHandler.UserHandler{UserUseCase: mockUseCase}
	httpHandler.NewUserHttpHandler(e, mockUseCase)

	reqBody, err := json.Marshal(userDTO)
	assert.NoError(t, err)

	req := httptest.NewRequest(echo.POST, "/users", strings.NewReader(string(reqBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	// Call the handler
	if assert.NoError(t, handler.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var userResponse dto.UserResponseDTO
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &userResponse)) {
			assert.Equal(t, userID, userResponse.ID)
			assert.Equal(t, userDTO.Name, userResponse.Name)
			assert.Equal(t, userDTO.Email, userResponse.Email)
			assert.Equal(t, "xxx", userResponse.Password)
		}
	}

	mockUseCase.AssertExpectations(t)
}
