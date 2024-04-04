package service

import (
	"database/sql"
	"fmt"

	"github.com/molinavtomas/labora-api-personas/db_"
	"github.com/molinavtomas/labora-api-personas/models"
)

func ObtenerPersona(db *sql.DB, id int) (models.PersonaExtendida, error) {
	persona, err := db_.ObtenerPersonaDB(db, id)
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

func ModificarPersona(db *sql.DB, p models.Persona) (models.Persona, error) {
	personaAux, err := db_.ObtenerPersonaDB(db, p.ID)
	if err != nil {
		return models.Persona{}, err
	}

	p, err = db_.ModificarPersonaDB(db, p, personaAux)
	if err != nil {
		return models.Persona{}, err
	}

	fmt.Println("Datos obtenidos correctamente")
	return p, nil

}
