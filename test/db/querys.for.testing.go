package test

import (
	"database/sql"
	"fmt"

	"github.com/molinavtomas/labora-api-personas/db_"
	"github.com/molinavtomas/labora-api-personas/models"
)

func ModificarPersonaForTesting(db *sql.DB, p models.Persona, personaAux models.Persona) (models.Persona, error) {

	personaAux, err := db_.ObtenerPersonaDB(db, p.ID)
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

	query := "UPDATE personas SET nombre = $1, apellido = $2, edad = $3, country_code = $4 WHERE id = $5 RETURNING *;"
	row := db.QueryRow(query, personaAux.Nombre, personaAux.Apellido, personaAux.Edad, personaAux.CountryCode, personaAux.ID)

	if err := row.Scan(&p.ID, &p.Nombre, &p.Apellido, &p.Edad, &p.CountryCode); err != nil {
		return models.Persona{}, fmt.Errorf("error al devolver persona actualizada, error: %w", err)
	}

	fmt.Println("Persona modificada correctamente")
	return p, nil

}

func CreatePersonaTesting(db *sql.DB, p models.Persona) (int, error) {
	if !p.Validate() {
		return -1, &models.ErrorPersonaInvalida{Mensaje: "La persona no es válida"}
	}

	nombre := p.Nombre
	apellido := p.Apellido
	edad := p.Edad
	country_code := p.CountryCode

	// Preparar la consulta SQL
	query := "INSERT INTO personas (nombre, apellido, edad, country_code) VALUES ($1, $2, $3, $4) RETURNING id;"

	// Ejecutar la consulta
	row := db.QueryRow(query, nombre, apellido, edad, country_code)

	var id_ int
	if err := row.Scan(&id_); err != nil {
		return -1, fmt.Errorf("error al crear persona en la base de datos: %w", err)
	}

	fmt.Println("Persona creada correctamente")
	return id_, nil
}
