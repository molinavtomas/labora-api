package routes

import (
	"encoding/json"
	"net/http"

	"github.com/molinavtomas/labora-api-personas/db_"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db_.ConectionDB()

	personas, err := db_.ObtenerPersonas(db)
	if err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Convertir el slice de Personas a formato JSON
	personasJSON, err := json.Marshal(personas)
	if err != nil {
		http.Error(w, "ERROR al convertir a JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Establecer encabezado Content-Type y escribir respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(personasJSON)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola putos"))
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
}
