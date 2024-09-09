package configs_test

import (
	"os"
	"testing"

	"github.com/quippv/go-clean/configs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EnvLookupTestSuite struct {
	suite.Suite
}

func (suite *EnvLookupTestSuite) SetupTest() {
	os.Setenv("TEST_ENV_VAR", "test_value")
}

func (suite *EnvLookupTestSuite) TearDownTest() {
	os.Unsetenv("TEST_ENV_VAR")
}

func TestEnvLookup(t *testing.T) {
	suite.Run(t, new(EnvLookupTestSuite))
}

func (suite *EnvLookupTestSuite) TestEnvLookupSuccess() {
	value := configs.EnvLookup("TEST_ENV_VAR")
	assert.Equal(suite.T(), "test_value", value)
}

func (suite *EnvLookupTestSuite) TestEnvLookupPanic() {
	assert.PanicsWithValue(suite.T(), "[ENV CONFIG]: MISSING_ENV not found in environment variable", func() {
		configs.EnvLookup("MISSING_ENV")
	})
}
