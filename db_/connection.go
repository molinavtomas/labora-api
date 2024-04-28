package db_

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DBConnection *sql.DB

func ConectionDB() error {
	//Leer variables de entorno
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	//Usar las variables de entorno en la cadena de conexión
	connectionString := fmt.Sprintf("user=%s dbname=%s host=%s sslmode=%s password=%s port=%s", dbUser, dbName, dbHost, "disable", dbPassword, dbPort)
	var err error
	DBConnection, err = sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("error al abrir la conexión a la base de datos: %w", err)
	}

	// Verificar si la conexión se estableció correctamente
	if err = DBConnection.Ping(); err != nil {
		return fmt.Errorf("error al establecer la conexión a la base de datos: %w", err)
	}

	fmt.Println("Conexión exitosa a la base de datos!")
	return nil
}
