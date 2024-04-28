package test

import (
	"database/sql"
	"fmt"

	"github.com/molinavtomas/labora-api-personas/db_"
	"github.com/molinavtomas/labora-api-personas/models"
)

func CreatePersonaForTesting(db *sql.DB, p models.Persona) (int, error) {
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

func ObtenerPersonasForTesting(db *sql.DB) ([]models.Persona, error) {
	// Preparar la consulta SQL
	query := "SELECT * FROM personas"

	// Ejecutar la consulta SQL
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener las personas en la base de datos: %w", err)
	}

	defer rows.Close()

	// Iterar sobre los resultados y mapearlos a la estructura Persona
	var personas []models.Persona

	for rows.Next() {
		var persona models.Persona
		if err := rows.Scan(&persona.ID, &persona.Nombre, &persona.Apellido, &persona.Edad, &persona.CountryCode); err != nil {
			return nil, fmt.Errorf("error al escanear fila: %w", err)
		}
		personas = append(personas, persona)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error en las filas: %w", err)
	}

	fmt.Println("Datos obtenidos correctamente")
	return personas, nil
}

func ObtenerPersonaDBForTesting(db *sql.DB, id int) (models.Persona, error) {
	// Preparar la consulta SQL
	query := "SELECT * FROM personas where id = $1"

	// Ejecutar la consulta SQL
	row := db.QueryRow(query, id)

	var persona models.Persona
	if err := row.Scan(&persona.ID, &persona.Nombre, &persona.Apellido, &persona.Edad, &persona.CountryCode); err != nil {
		return models.Persona{}, fmt.Errorf("persona con ID %d no encontrada en la base de datos, error: %w", id, err)
	}

	fmt.Println("Datos obtenidos correctamente")
	return persona, nil
}

func ModificarPersonaForTesting(db *sql.DB, p models.Persona, personaAux models.Persona) (models.Persona, error) {

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

	query := "UPDATE personas SET nombre = $1, apellido = $2, edad = $3, country_code = $4 WHERE id = $5 RETURNING *;"
	row := db.QueryRow(query, personaAux.Nombre, personaAux.Apellido, personaAux.Edad, personaAux.CountryCode, personaAux.ID)

	if err := row.Scan(&p.ID, &p.Nombre, &p.Apellido, &p.Edad, &p.CountryCode); err != nil {
		return models.Persona{}, fmt.Errorf("error al devolver persona actualizada, error: %w", err)
	}

	fmt.Println("Persona modificada correctamente")
	return p, nil

}

func EliminarPersonaDBForTesting(db *sql.DB, id int) error {
	// Preparar la consulta SQL de inserción
	query := "DELETE FROM personas WHERE id = $1 RETURNING id;"

	// Ejecutar la consulta SQL de inserción
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar persona en la base de datos: %w", err)
	}

	fmt.Println("Persona eliminada correctamente")
	return nil
}
