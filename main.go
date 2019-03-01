package main

import (
	"log"
	"net/http"

	"github.com/vickleford/motn/api/v1"
)

func main() {
	apiv1 := v1.ServeV1()

	server := http.NewServeMux()
	server.Handle("/v1/", apiv1)

	err := http.ListenAndServe(":8080", server)
	log.Fatal(err)
}
