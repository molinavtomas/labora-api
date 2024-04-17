package test

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//Copia de la funcion ConectionDB con una base de datos hechas para pruebas.

func ConectionDBTesting(conectionString string) (*sql.DB, error) {
	// Configurar la conexión a PostgreSQL
	db, err := sql.Open("postgres", conectionString)
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
