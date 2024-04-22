package main

import (
	"fmt"
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

	fmt.Println("Servidor iniciado en el puerto 8080...")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
