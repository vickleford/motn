package motn

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func fakeNow() time.Time {
	return time.Unix(1550090938, 0)
}

func TestMotn(t *testing.T) {
	motnHandler := new(Motn)
	motnHandler.NowFunc = fakeNow

	r, _ := http.NewRequest("GET", "/v1/motn", nil)
	w := httptest.NewRecorder()
	motnHandler.ServeHTTP(w, r)

	if expected := 200; w.Code != expected {
		t.Errorf("Expected response code %d, got %d", expected, w.Code)
	}

	jsonresp := new(response)
	err := json.Unmarshal(w.Body.Bytes(), jsonresp)
	if err != nil {
		t.Errorf("Couldn't parse that json: %s", err.Error())
	}

	if expected := "Code for the masses"; jsonresp.Message != expected {
		t.Errorf("Expected message to be '%s', got %s", expected, jsonresp.Message)
	}

	if expected := int64(1550090938); jsonresp.Timestamp != expected {
		t.Errorf("Expected timestamp to be %d, got %d", expected, jsonresp.Timestamp)
	}
}
