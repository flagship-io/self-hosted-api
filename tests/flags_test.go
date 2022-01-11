package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlagsRoute(t *testing.T) {
	w := httptest.NewRecorder()
	jsonStr := []byte(`{"visitor_id": "test"}`)
	req, _ := http.NewRequest("POST", "/v2/flags", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), `{"`)
	jsonResult := map[string]map[string]interface{}{}
	json.Unmarshal(w.Body.Bytes(), &jsonResult)
	assert.Greater(t, len(jsonResult), 0)

	for _, v := range jsonResult {
		assert.NotNil(t, v["metadata"])
	}
}
