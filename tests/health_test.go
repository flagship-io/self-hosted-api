package runner

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flagship-io/self-hosted-api/runner"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestActivateRoute(t *testing.T) {
	viper.SetDefault("env_id", "blvo2kijq6pg023l8edg")
	viper.SetDefault("api_key", "fake-key")

	router, err := runner.InitRouter(runner.GetOptionsFromConfig())
	assert.NotNil(t, router)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v2/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"status":"READY"}`, w.Body.String())
}
