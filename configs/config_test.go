package configs_test

import (
	"os"
	"testing"

	"github.com/quippv/go-clean/configs"
	"github.com/stretchr/testify/assert"
)

func TestEnvLoad(t *testing.T) {
	testEnvFile := ".env.test"

	err := os.WriteFile(testEnvFile, []byte("TEST_ENV_VAR=hello\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test .env file: %v", err)
	}
	defer os.Remove(testEnvFile)

	configs.EnvLoad(testEnvFile)

	value, exists := os.LookupEnv("TEST_ENV_VAR")
	assert.True(t, exists, "Environment variable should be set")
	assert.Equal(t, "hello", value, "Environment variable value should match")
}
