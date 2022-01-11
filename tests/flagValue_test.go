package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlagValueRoute(t *testing.T) {
	w := httptest.NewRecorder()
	jsonStr := []byte(`{"visitor_id": "test"}`)
	req, _ := http.NewRequest("POST", "/v2/flags", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), `{"`)
	jsonResult := map[string]map[string]interface{}{}
	json.Unmarshal(w.Body.Bytes(), &jsonResult)
	assert.Greater(t, len(jsonResult), 0)

	var firstKeyName string
	var firstKeyValue map[string]interface{}
	for k, v := range jsonResult {
		firstKeyName = k
		firstKeyValue = v
		break
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/v2/flags/"+firstKeyName+"/value", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var value interface{}
	json.Unmarshal(w.Body.Bytes(), &value)
	assert.Equal(t, firstKeyValue["value"], value)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/v2/flags/not_exists/value", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "{\"error\":\"Flag not_exists not found\"}", w.Body.String())
}
