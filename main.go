package main

import (
	"net/http"

	"github.com/Miqbal20/golang-restapi/api"
)

func main() {
	http.HandleFunc("/", api.Handler)
	http.ListenAndServe(":8080", nil)
}
