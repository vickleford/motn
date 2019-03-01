package healthz

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthz(t *testing.T) {
	healthzHandler := new(Healthz)
	r, _ := http.NewRequest("GET", "/v1/healthz", nil)
	w := httptest.NewRecorder()
	healthzHandler.ServeHTTP(w, r)

	if expected := 200; w.Code != expected {
		t.Errorf("Expected response code %d, got %d", expected, w.Code)
	}

	jsonresp := new(response)
	err := json.Unmarshal(w.Body.Bytes(), jsonresp)
	if err != nil {
		t.Errorf("Couldn't parse that json: %s", err.Error())
	}

	if expected := "green"; jsonresp.Status != expected {
		t.Errorf("Expected status to be %s, got '%s'", expected, jsonresp.Status)
	}
}
