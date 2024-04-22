package db_

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConectionDB() (*sql.DB, error) {
	// Leer variables de entorno
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Usar las variables de entorno en la cadena de conexión
	connectionString := fmt.Sprintf("user=%s dbname=%s host=%s sslmode=%s password=%s port=%s", dbUser, dbName, dbHost, "disable", dbPassword, dbPort)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error al abrir la conexión a la base de datos: %w", err)
	}

	// Verificar si la conexión se estableció correctamente
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error al establecer la conexión a la base de datos: %w", err)
	}

	// Ejecutar la consulta para crear la tabla personas si no existe
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS personas (
		id SERIAL PRIMARY KEY,
		nombre VARCHAR(100) NOT NULL,
		apellido VARCHAR(100) NOT NULL,
		edad INTEGER NOT NULL,
		country_code VARCHAR(10) NOT NULL
	);`)
	if err != nil {
		return nil, fmt.Errorf("error al crear la tabla personas: %w", err)
	}

	fmt.Println("Conexión exitosa")
	return db, nil
}
