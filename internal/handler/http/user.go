package http

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/quippv/go-clean/internal/dto"
	"github.com/quippv/go-clean/internal/entity"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
	Reason  string `json:"reason"`
}

// defines the interface for user repository contract
type UserUseCase interface {
	GetUser(id uuid.UUID) (*entity.User, error)
	RegisterUser(user *entity.User) (uuid.UUID, error)
}

type UserHandler struct {
	UserUseCase UserUseCase
}

// registers the user routes
func NewUserHttpHandler(e *echo.Echo, userUseCase UserUseCase) {
	handler := &UserHandler{userUseCase}
	e.GET("/users/:id", handler.GetUser)
	e.POST("/users", handler.CreateUser)
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Get details of a user by their UUID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} dto.UserResponseDTO
// @Failure 400 {object} ResponseError "Invalid UUID provided"
// @Failure 404 {object} ResponseError "User not found"
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{
			Message: err.Error(),
			Reason:  "The provided ID is not a valid UUID.",
		})
	}

	user, err := h.UserUseCase.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, ResponseError{
			Message: err.Error(),
			Reason:  "The user with the provided ID does not exist.",
		})
	}

	userResponseDto := dto.ToDTO(user)

	return c.JSON(http.StatusOK, userResponseDto)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body dto.CreateUserDTO true "Create User"
// @Success 201 {object} dto.UserResponseDTO
// @Failure 400 {object} ResponseError "Invalid request payload"
// @Failure 500 {object} ResponseError "Internal server error"
// @Router /users [post]
func (h *UserHandler) CreateUser(c echo.Context) error {
	var userDTO dto.CreateUserDTO
	if err := c.Bind(&userDTO); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{
			Message: "Invalid request payload",
			Reason:  err.Error(),
		})
	}

	user := userDTO.ToEntity()

	id, err := h.UserUseCase.RegisterUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{
			Message: err.Error(),
			Reason:  err.Error(),
		})
	}

	user.ID = id
	userResponseDto := dto.ToDTO(user)

	return c.JSON(http.StatusCreated, userResponseDto)
}
