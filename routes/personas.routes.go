package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/molinavtomas/labora-api-personas/models"
	"github.com/molinavtomas/labora-api-personas/service"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	personas, err := service.ObtenerPersonas()
	if err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(personas) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{}"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(personas); err != nil {
		http.Error(w, "ERROR al convertir a JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	idAsInt, _ := strconv.Atoi(idString)

	persona, err := service.ObtenerPersona(idAsInt)
	if err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(persona); err != nil {
		http.Error(w, "ERROR al convertir a JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var persona models.Persona
	if err := decoder.Decode(&persona); err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err := service.CrearPersona(&persona)
	if err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(persona); err != nil {
		http.Error(w, "ERROR al convertir a JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

}

func PutUserHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var persona models.Persona
	if err := decoder.Decode(&persona); err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusBadRequest)
		return
	}

	var err error
	if persona, err = service.ModificarPersona(persona); err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(persona); err != nil {
		http.Error(w, "ERROR al convertir a JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	idAsInt, _ := strconv.Atoi(idString)

	err := service.EliminarPersona(idAsInt)
	if err != nil {
		http.Error(w, "ERROR: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Persona eliminada correctamente"))
	w.WriteHeader(http.StatusOK)

}
