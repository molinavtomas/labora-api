package db_

import (
	"database/sql"
	"fmt"

	"github.com/molinavtomas/labora-api-personas/models"
)

func CreatePersona(db *sql.DB, p models.Persona) error {
	nombre := p.Nombre
	apellido := p.Apellido
	edad := p.Edad
	country_code := p.CountryCode

	// Preparar la consulta SQL
	query := "INSERT INTO personas (nombre, apellido, edad, country_code) VALUES ($1, $2, $3, $4)"

	// Ejecutar la consulta
	_, err := db.Exec(query, nombre, apellido, edad, country_code)
	if err != nil {
		return fmt.Errorf("error al crear persona en la base de datos: %w", err)
	}

	fmt.Println("Persona creada correctamente")
	return nil
}

func ObtenerPersonas(db *sql.DB) ([]models.Persona, error) {
	// Preparar la consulta SQL
	query := "SELECT * FROM personas"

	// Ejecutar la consulta SQL de inserción
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error al obtener las personas en la base de datos: %w", err)
	}

	defer rows.Close()

	// Iterar sobre los resultados y mapearlos a la estructura Persona
	var personas []models.Persona

	for rows.Next() {
		var persona models.Persona
		if err := rows.Scan(&persona.ID, &persona.Nombre, &persona.Apellido, &persona.Edad, &persona.CountryCode); err != nil {
			return nil, fmt.Errorf("Error al escanear fila: %w", err)
		}
		personas = append(personas, persona)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error en las filas: %w", err)
	}

	fmt.Println("Datos obtenidos correctamente")
	return personas, nil
}

func ObtenerPersona(db *sql.DB, id int) error {
	// Preparar la consulta SQL
	query := "SELECT * FROM personas where id = $1"

	// Ejecutar la consulta SQL de inserción
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Error al obtener la persona en la base de datos: %w", err)
	}

	fmt.Println("Datos obtenidos correctamente")
	return nil
}

func ModificarNombrePersona(db *sql.DB, id int, p models.Persona, nuevo_nombre string) error {
	// Preparar la consulta SQL de modificación
	query := "UPDATE personas SET nombre = $2 WHERE id = $1"
	p.Nombre = nuevo_nombre

	// Ejecutar la consulta SQL de modificación
	_, err := db.Exec(query, id, nuevo_nombre)
	if err != nil {
		return fmt.Errorf("error al modificar persona en la base de datos: %w", err)
	}

	fmt.Println("Persona modificada correctamente")
	return nil
}

func ModificarApellidoPersona(db *sql.DB, id int, p models.Persona, nuevo_apellido string) error {
	// Preparar la consulta SQL de modificación
	query := "UPDATE personas SET apellido = $2 WHERE id = $1"
	p.Apellido = nuevo_apellido

	// Ejecutar la consulta SQL de modificación
	_, err := db.Exec(query, id, nuevo_apellido)
	if err != nil {
		return fmt.Errorf("error al modificar persona en la base de datos: %w", err)
	}

	fmt.Println("Persona modificada correctamente")
	return nil
}

func ModificarEdadPersona(db *sql.DB, id int, p models.Persona, nueva_edad int) error {
	// Preparar la consulta SQL de modificación
	query := "UPDATE personas SET edad = $2 WHERE id = $1"
	p.Edad = nueva_edad

	// Ejecutar la consulta SQL de modificación
	_, err := db.Exec(query, id, nueva_edad)
	if err != nil {
		return fmt.Errorf("error al modificar persona en la base de datos: %w", err)
	}

	fmt.Println("Persona modificada correctamente")
	return nil
}

func ModificarCountryCodePersona(db *sql.DB, id int, p models.Persona, nuevo_country_code string) error {
	// Preparar la consulta SQL de modificación
	query := "UPDATE personas SET country_code = $2 WHERE id = $1"
	p.CountryCode = nuevo_country_code

	// Ejecutar la consulta SQL de modificación
	_, err := db.Exec(query, id, nuevo_country_code)
	if err != nil {
		return fmt.Errorf("error al modificar persona en la base de datos: %w", err)
	}

	fmt.Println("Persona modificada correctamente")
	return nil
}

func EliminarPersona(db *sql.DB, id int) error {
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
