package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	httpHandler "github.com/quippv/go-clean/internal/handler/http"
	"github.com/stretchr/testify/assert"
	"github.com/swaggo/echo-swagger"
)

func TestNewSwaggerHttpHandler(t *testing.T) {
	e := echo.New()

	httpHandler.NewSwaggerHttpHandler(e)

	req := httptest.NewRequest(http.MethodGet, "/docs/index.html", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, echoSwagger.WrapHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
