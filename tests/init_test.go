package tests

import (
	"os"
	"testing"

	"github.com/flagship-io/self-hosted-api/runner"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	viper.SetDefault("env_id", "blvo2kijq6pg023l8edg")
	viper.SetDefault("api_key", "fake-key")

	router, _ = runner.InitRouter(runner.GetOptionsFromConfig())
	code := m.Run()
	os.Exit(code)
}
