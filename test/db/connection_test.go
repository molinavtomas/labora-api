package test

import (
	"testing"
)

func TestConectionDB_Successful(t *testing.T) {
	// Conectar a la base de datos
	conectionString := "user=postgres dbname=personas_db_testing host=localhost sslmode=disable password=kys0128tomas port=5432"
	db, err := ConectionDBTesting(conectionString)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	defer db.Close()
}

func TestConectionDB_OpenError(t *testing.T) {
	//sql.Open solo verifica si la cadena de conexión está bien formada, pero no realiza una conexión real a la base de datos en ese momento.
	//Por lo que para combrobar que la coneccion es correcta debo hacer un ping, sino no devolveria ningun error
	conectionString := "FALSE"

	db, err := ConectionDBTesting(conectionString)
	// Verificar si la función devuelve un error
	if err == nil {
		t.Error("Se esperaba un error al abrir la conexión, pero no se recibió ninguno")
	}
	// Verificar si db es nulo
	if db != nil {
		t.Error("Se esperaba que db sea nulo si la conexión falla, pero no lo es")
	}

	t.Logf("Prueba finalizada. No se pudo abrir la base de datos.")
}

func TestConectionDB_PingError(t *testing.T) {
	//Modifique el puerto de la base de datos
	conectionString := "user=postgres dbname=personas_db_testing host=localhost sslmode=disable password=kys0128tomas port=5434"
	db, err := ConectionDBTesting(conectionString)
	// Verificar si la función devuelve un error
	if err == nil {
		t.Error("Se esperaba un error al abrir la conexión, pero no se recibió ninguno")
	}
	// Verificar si db es nulo
	if db != nil {
		t.Error("Se esperaba que db sea nulo si la conexión falla, pero no lo es")
	}

	t.Logf("Prueba finalizada. No se pudo concrear el Ping a la base de datos.")
}
