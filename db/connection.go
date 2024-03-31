package db

import (
	"database/sql"
	"fmt"
)

func conectionDB() (*sql.DB, error) {
	// Configurar la conexión a PostgreSQL
	db, err := sql.Open("postgres", "user=postgres dbname=personas_db host=localhost sslmode=disable password=kys0128tomas port=5433")
	if err != nil {
		return nil, fmt.Errorf("error al abrir la conexión a la base de datos: %w", err)
	}

	// Verificar si la conexión se estableció correctamente
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error al establecer la conexión a la base de datos: %w", err)
	}

	fmt.Println("Conexión exitosa")
	return db, nil
}
