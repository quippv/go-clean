package usecase

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	passwordUseCase := NewPasswordUseCase()

	tests := []struct {
		password string
	}{
		{"password123"},
		{"anotherPassword!@#"},
		{"short"},
	}

	for _, test := range tests {
		t.Run(test.password, func(t *testing.T) {
			hashedPassword, err := passwordUseCase.HashPassword(test.password)
			assert.NoError(t, err)
			assert.NotEmpty(t, hashedPassword)
			parts := strings.Split(hashedPassword, "$")
			assert.Len(t, parts, 2)
		})
	}
}
