package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/molinavtomas/labora-api-personas/routes"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", router)
}
