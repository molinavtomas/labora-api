package service

import (
	"fmt"

	"github.com/molinavtomas/labora-api-personas/db_"
	"github.com/molinavtomas/labora-api-personas/models"
)

func CrearPersona(p *models.Persona) (int, error) {
	if !p.Validate() {
		return -1, &models.ErrorPersonaInvalida{Mensaje: "La persona no es válida"}
	}

	id, err := db_.CreatePersona(*p)
	if err != nil {
		return -1, err
	}

	//Actualizamos ID de la nueva persona
	p.ID = id

	return id, nil

}

func ObtenerPersona(id int) (models.PersonaExtendida, error) {
	persona, err := db_.ObtenerPersonaDB(id)
	if err != nil {
		return models.PersonaExtendida{}, err
	}

	countryInfo, err := getCountryInfo(persona.CountryCode)
	if err != nil {
		return models.PersonaExtendida{}, fmt.Errorf("error al obtener información del país %v", err)
	}

	fmt.Println("Datos obtenidos correctamente")

	return models.PersonaExtendida{
		Persona:     persona,
		CountryInfo: countryInfo,
	}, nil
}

func ObtenerPersonas() ([]models.Persona, error) {
	personas, err := db_.ObtenerPersonas()
	if err != nil {
		return nil, fmt.Errorf("error al obtener personas en el servicio: %w", err)
	}

	return personas, nil
}

func ModificarPersona(p models.Persona) (models.Persona, error) {
	personaAux, err := db_.ObtenerPersonaDB(p.ID)
	if err != nil {
		return models.Persona{}, err
	}
	if p.ID != personaAux.ID {
		return models.Persona{}, &models.ErrorPersonaInvalida{Mensaje: "La persona auxiliar no es valida"}
	}
	if p.Nombre != "" {
		personaAux.Nombre = p.Nombre
	}
	if p.Apellido != "" {
		personaAux.Apellido = p.Apellido
	}
	if p.Edad > 0 {
		personaAux.Edad = p.Edad
	}
	if p.CountryCode != "" {
		personaAux.CountryCode = p.CountryCode
	}

	p, err = db_.ModificarPersonaDB(p, personaAux)
	if err != nil {
		return models.Persona{}, err
	}

	fmt.Println("Datos obtenidos correctamente")
	return p, nil

}

func EliminarPersona(id int) error {
	err := db_.EliminarPersonaDB(id)
	if err != nil {
		return err
	}

	return nil
}
