package v1

import (
	"net/http"
	"time"

	"github.com/vickleford/motn/api/v1/healthz"
	"github.com/vickleford/motn/api/v1/motn"
)

func ServeV1() *http.ServeMux {
	h := http.NewServeMux()

	motn := new(motn.Motn)
	motn.NowFunc = time.Now

	healthz := new(healthz.Healthz)

	h.Handle("/v1/motn", motn)
	h.Handle("/v1/healthz", healthz)

	return h
}
