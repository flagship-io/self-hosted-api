package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlagRoute(t *testing.T) {
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
	req, _ = http.NewRequest("POST", "/v2/flags/"+firstKeyName, bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	jsonResultKey := map[string]interface{}{}
	log.Println(w.Body.String())
	err := json.Unmarshal(w.Body.Bytes(), &jsonResultKey)
	log.Println(err)
	assert.Greater(t, len(jsonResultKey), 0)
	assert.Equal(t, firstKeyValue, jsonResultKey)
}
