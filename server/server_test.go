package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testServe = CreateServer()

func TestPostRouteMiss(t *testing.T) {

	reqBody := Secret{Key: "testkey"}
	reqBodyBytes, _ := json.Marshal(reqBody)

	respBody := Response{Content: ""}
	respBodyBytes, _ := json.Marshal(respBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/secret", bytes.NewBuffer(reqBodyBytes))
	testServe.Router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, respBodyBytes, w.Body.Bytes())
}

func TestPostRouteHit(t *testing.T) {
	testServe.controller.UpdateSecret("testkey", "testcontent")

	reqBody := Secret{Key: "testkey"}
	reqBodyBytes, _ := json.Marshal(reqBody)

	respBody := Response{Content: "testcontent"}
	respBodyBytes, _ := json.Marshal(respBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/secret", bytes.NewBuffer(reqBodyBytes))
	testServe.Router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, respBodyBytes, w.Body.Bytes())
}

func TestPutRoute(t *testing.T) {
	testServe.controller.UpdateSecret("testkey", "testcontent")

	reqBody := Secret{Key: "testkey", Content: "differenttestcontent"}
	reqBodyBytes, _ := json.Marshal(reqBody)

	respBody := Response{Content: "differenttestcontent"}
	respBodyBytes, _ := json.Marshal(respBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/secret", bytes.NewBuffer(reqBodyBytes))
	testServe.Router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, respBodyBytes, w.Body.Bytes())
}

func TestDeleteRoute(t *testing.T) {
	testServe.controller.UpdateSecret("testkey", "testcontent")

	reqBody := Secret{Key: "testkey"}
	reqBodyBytes, _ := json.Marshal(reqBody)

	respBody := Response{Content: ""}
	respBodyBytes, _ := json.Marshal(respBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/secret", bytes.NewBuffer(reqBodyBytes))
	testServe.Router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, respBodyBytes, w.Body.Bytes())
}
