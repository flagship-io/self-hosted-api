package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCampaignRoute(t *testing.T) {
	w := httptest.NewRecorder()
	jsonStr := []byte(`{"visitor_id": "test"}`)
	req, _ := http.NewRequest("POST", "/v2/campaigns/not_found", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Contains(t, w.Body.String(), `not_found`)
}
