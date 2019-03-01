package motn

import (
	"encoding/json"
	"net/http"
	"time"
)

type response struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

type Motn struct {
	NowFunc func() time.Time
}

func (h *Motn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := response{
		Message:   "Code for the masses",
		Timestamp: h.NowFunc().Unix(),
	}

	e := json.NewEncoder(w)
	e.Encode(resp)
}
