package main

import (
	"net/http"

	"github.com/molinavtomas/labora-api-personas/routes"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("GET /personas", routes.GetUsersHandler)
	router.HandleFunc("GET /personas/{id}", routes.GetUserHandler)
	router.HandleFunc("POST /personas", routes.PostUserHandler)
	router.HandleFunc("PUT /personas", routes.PutUserHandler)
	router.HandleFunc("DELETE /personas/{id}", routes.DeleteUserHandler)

	http.ListenAndServe(":3000", router)
}
