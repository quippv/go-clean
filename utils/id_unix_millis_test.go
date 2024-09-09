package utils_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/quippv/go-clean/utils"
	"github.com/stretchr/testify/assert"
)

func TestGenerateIDAndUnixMillis(t *testing.T) {
	id, currentMillis, returnedMillis := utils.GenerateIDAndUnixMillis()

	_, err := uuid.Parse(id.String())
	assert.NoError(t, err, "generated UUID should be valid")

	assert.True(t, returnedMillis >= currentMillis-1000, "returnedMillis should be close to currentMillis")
	assert.True(t, returnedMillis <= currentMillis+1000, "returnedMillis should be close to currentMillis")
}
