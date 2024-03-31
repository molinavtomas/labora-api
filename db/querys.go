package db

import (
	"database/sql"
	"fmt"
)

func createPersona(db *sql.DB, nombre string, apellido string, edad int, country_code string) error {
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

func obtenerPersonas(db *sql.DB) error {
	// Preparar la consulta SQL
	query := "SELECT * FROM personas"

	// Ejecutar la consulta SQL de inserción
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("Error al obtener las personas en la base de datos: %w", err)
	}

	fmt.Println("Datos obtenidos correctamente")
	return nil
}

func obtenerPersona(db *sql.DB, id int) error {
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

func modificarPersona(db *sql.DB, id int, nombre string, apellido string, edad int, country_code string) error {
	// Preparar la consulta SQL de modificación
	query := "UPDATE personas SET nombre = $2, apellido = $3, edad = $4, country_code = $5 WHERE id = $1"

	// Ejecutar la consulta SQL de modificación
	_, err := db.Exec(query, id, nombre, apellido, edad, country_code)
	if err != nil {
		return fmt.Errorf("error al modificar persona en la base de datos: %w", err)
	}

	fmt.Println("Persona modificada correctamente")
	return nil
}

func eliminarPersona(db *sql.DB, id int) error {
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
