package healthz

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status string `json:"status"`
}

type Healthz struct{}

func (h *Healthz) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := response{Status: "green"}

	e := json.NewEncoder(w)
	e.Encode(resp)
}
