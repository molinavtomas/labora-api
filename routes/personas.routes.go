package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/molinavtomas/labora-api-personas/db_"
	"github.com/molinavtomas/labora-api-personas/models"
	"github.com/molinavtomas/labora-api-personas/service"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db_.ConectionDB()
	if err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusInternalServerError)
		return
	}

	personas, err := service.ObtenerPersonas(db)
	if err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(personas); err != nil {
		http.Error(w, "ERROR al convertir a JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {

	db, err := db_.ConectionDB()
	if err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusInternalServerError)
		return
	}

	idString := r.PathValue("id")
	idAsInt, _ := strconv.Atoi(idString)

	persona, err := service.ObtenerPersona(db, idAsInt)
	if err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(persona); err != nil {
		http.Error(w, "ERROR al convertir a JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db_.ConectionDB()
	if err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var persona models.Persona
	if err = decoder.Decode(&persona); err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = service.CrearPersona(db, &persona)
	if err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(persona); err != nil {
		http.Error(w, "ERROR al convertir a JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

}

func PutUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db_.ConectionDB()
	if err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var persona models.Persona
	if err = decoder.Decode(&persona); err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusBadRequest)
		return
	}

	if persona, err = service.ModificarPersona(db, persona); err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(persona); err != nil {
		http.Error(w, "ERROR al convertir a JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db_.ConectionDB()
	if err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusInternalServerError)
		return
	}

	idString := r.PathValue("id")
	idAsInt, _ := strconv.Atoi(idString)

	err = service.EliminarPersona(db, idAsInt)
	if err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Persona eliminada correctamente")
	w.WriteHeader(http.StatusOK)

}
