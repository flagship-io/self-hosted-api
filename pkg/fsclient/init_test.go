package fsclient

import (
	"os"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestInitFsClient(t *testing.T) {
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	fsClient, err := InitFsClient()

	assert.NotNil(t, err)
	assert.Nil(t, fsClient)
	assert.Contains(t, err.Error(), "EnvID is required")

	os.Setenv("ENV_ID", "test_env_id")
	fsClient, err = InitFsClient()

	assert.NotNil(t, err)
	assert.Nil(t, fsClient)
	assert.Contains(t, err.Error(), "APIKey is required")

	os.Setenv("API_KEY", "test_api_key")
	fsClient, err = InitFsClient()

	assert.NotNil(t, fsClient)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "403 Forbidden")
}
