package configs

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/quippv/go-clean/internal/handler/http"
	"github.com/quippv/go-clean/internal/repository/postgres"
	"github.com/quippv/go-clean/internal/usecase"
)

func ComposeAppSymphony(dbConn *sql.DB, e *echo.Echo) {
	// Prepare Repository
	userRepo := postgres.NewUserRepository(dbConn)

	// Build service layer
	password := usecase.NewPasswordUseCase()
	user := usecase.NewUserUseCase(userRepo, password)

	// Initialize HTTP handlers
	http.NewSwaggerHttpHandler(e)
	http.NewUserHttpHandler(e, user)
}
