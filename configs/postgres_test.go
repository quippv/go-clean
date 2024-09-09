package configs_test

import (
	"testing"

	"github.com/quippv/go-clean/configs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpen(t *testing.T) {

	// Load the environment variables
	configs.EnvLoad("../.env")

	// Get default postgres config
	config := configs.DefaultPostgresConfig()

	// Open the database connection
	db, err := configs.Open(config)
	require.NoError(t, err)
	require.NotNil(t, db)
	defer db.Close()

	// Test if the connection is established by pinging the database
	err = db.Ping()
	assert.NoError(t, err)
}
