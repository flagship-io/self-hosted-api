package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCampaignsRoute(t *testing.T) {
	w := httptest.NewRecorder()
	jsonStr := []byte(`{"visitor_id": "test"}`)
	req, _ := http.NewRequest("POST", "/v2/campaigns", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), `"visitor_id":"test"`)
	assert.Contains(t, w.Body.String(), `"campaigns":[`)
}
