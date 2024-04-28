package db_

import (
	"fmt"

	"github.com/molinavtomas/labora-api-personas/models"
)

func CreatePersona(p models.Persona) (int, error) {
	nombre := p.Nombre
	apellido := p.Apellido
	edad := p.Edad
	country_code := p.CountryCode

	// Preparar la consulta SQL
	query := "INSERT INTO personas (nombre, apellido, edad, country_code) VALUES ($1, $2, $3, $4) RETURNING id;"

	// Ejecutar la consulta
	row := DBConnection.QueryRow(query, nombre, apellido, edad, country_code)

	var id_ int
	if err := row.Scan(&id_); err != nil {
		return -1, fmt.Errorf("error al crear persona en la base de datos: %w", err)
	}

	fmt.Println("Persona creada correctamente")
	return id_, nil
}

func ObtenerPersonas() ([]models.Persona, error) {
	var personas []models.Persona

	// Preparar la consulta SQL
	query := "SELECT * FROM personas"

	// Ejecutar la consulta SQL
	rows, err := DBConnection.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener las personas en la base de datos: %w", err)
	}

	defer rows.Close()

	// Iterar sobre los resultados y mapearlos a la estructura Persona

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

func ObtenerPersonaDB(id int) (models.Persona, error) {
	// Preparar la consulta SQL
	query := "SELECT * FROM personas where id = $1"

	// Ejecutar la consulta SQL
	row := DBConnection.QueryRow(query, id)

	var persona models.Persona
	if err := row.Scan(&persona.ID, &persona.Nombre, &persona.Apellido, &persona.Edad, &persona.CountryCode); err != nil {
		return models.Persona{}, fmt.Errorf("persona con ID %d no encontrada en la base de datos, error: %w", id, err)
	}

	fmt.Println("Datos obtenidos correctamente")
	return persona, nil
}

func ModificarPersonaDB(p models.Persona, personaAux models.Persona) (models.Persona, error) {

	query := "UPDATE personas SET nombre = $1, apellido = $2, edad = $3, country_code = $4 WHERE id = $5 RETURNING *;"
	row := DBConnection.QueryRow(query, personaAux.Nombre, personaAux.Apellido, personaAux.Edad, personaAux.CountryCode, personaAux.ID)

	if err := row.Scan(&p.ID, &p.Nombre, &p.Apellido, &p.Edad, &p.CountryCode); err != nil {
		return models.Persona{}, fmt.Errorf("error al devolver persona actualizada, error: %w", err)
	}

	fmt.Println("Persona modificada correctamente")
	return p, nil

}

func EliminarPersonaDB(id int) error {
	// Preparar la consulta SQL de inserción
	query := "DELETE FROM personas WHERE id = $1 RETURNING id;"

	// Ejecutar la consulta SQL de inserción
	_, err := DBConnection.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar persona en la base de datos: %w", err)
	}

	fmt.Println("Persona eliminada correctamente")
	return nil
}
