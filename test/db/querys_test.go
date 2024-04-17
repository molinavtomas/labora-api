package test

import (
	"testing"

	"github.com/molinavtomas/labora-api-personas/db_"
	"github.com/molinavtomas/labora-api-personas/models"
)

func TestObtenerPersonaPorID(t *testing.T) {
	// Conectar a la base de datos
	conectionString := "user=postgres dbname=personas_db_testing host=localhost sslmode=disable password=kys0128tomas port=5432"
	db, err := ConectionDBTesting(conectionString)
	if err != nil {
		t.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Agregar una persona de prueba a la base de datos
	row := db.QueryRow("INSERT INTO personas (nombre, apellido, edad, country_code) VALUES ($1, $2, $3, $4) RETURNING id;", "Prueba", "Prueba", 30, "US")

	var id_ int
	if err := row.Scan(&id_); err != nil {
		t.Fatalf("error al crear persona en la base de datos: %v", err)
	}

	// Obtener la persona recién agregada por su ID
	persona, err := db_.ObtenerPersonaDB(db, id_)
	if err != nil {
		t.Fatalf("Error al obtener la persona por ID: %v", err)
	}

	// Verificar que los detalles de la persona sean correctos
	expectedPersona := models.Persona{ID: id_, Nombre: "Prueba", Apellido: "Prueba", Edad: 30, CountryCode: "US"}
	if persona != expectedPersona {
		t.Errorf("Persona obtenida incorrecta. Se esperaba %v pero se obtuvo %v", expectedPersona, persona)
	}

	// Eliminar la persona de prueba de la base de datos
	_, _ = db.Exec("DELETE FROM personas WHERE id = $1", id_)
}

func TestObtenerPersonaPorID_Falla(t *testing.T) {
	// Conectar a la base de datos
	conectionString := "user=postgres dbname=personas_db_testing host=localhost sslmode=disable password=kys0128tomas port=5432"
	db, err := ConectionDBTesting(conectionString)
	if err != nil {
		t.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Intentar obtener una persona con un ID que no existe en la base de datos
	idInexistente := 9999 // ID que no existe
	persona, err := db_.ObtenerPersonaDB(db, idInexistente)
	if err == nil {
		t.Fatalf("Se esperaba un error al obtener la persona por ID %d, pero no se recibió ninguno", idInexistente)
	}

	// Verificar que la persona obtenida es una estructura vacía
	var personaVacia models.Persona
	if persona != personaVacia {
		t.Errorf("Se esperaba una estructura de persona vacía pero se recibió %v", persona)
	}
}

func TestCreatePersona(t *testing.T) {
	// Conectar a la base de datos
	conectionString := "user=postgres dbname=personas_db_testing host=localhost sslmode=disable password=kys0128tomas port=5432"
	db, err := ConectionDBTesting(conectionString)
	if err != nil {
		t.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Crear una persona válida
	p := models.Persona{
		Nombre:      "Prueba",
		Apellido:    "Prueba",
		Edad:        30,
		CountryCode: "US",
	}

	// Intentar crear la persona en la base de datos
	id, err := db_.CreatePersona(db, p)
	if err != nil {
		t.Fatalf("Error al crear la persona en la base de datos: %v", err)
	}

	// Verificar que se devolvió un ID válido
	if id < 0 {
		t.Fatalf("ID de persona inválido: %d", id)
	}

}

func TestCreatePersona_Falla(t *testing.T) {
	// Conectar a la base de datos
	conectionString := "user=postgres dbname=personas_db_testing host=localhost sslmode=disable password=kys0128tomas port=5432"
	db, err := ConectionDBTesting(conectionString)
	if err != nil {
		t.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Crear una persona con datos inválidos
	personaInvalida := models.Persona{} // Datos inválidos

	id, err := CreatePersonaTesting(db, personaInvalida)
	if err == nil {
		t.Fatalf("Se esperaba un error al crear una persona con datos inválidos, pero no se recibió ninguno")
	}

	// Verificar que el ID devuelto sea -1
	if id != -1 {
		t.Errorf("Se esperaba que el ID devuelto fuera -1 pero se recibió %d", id)
	}
}

func TestObtenerPersonas_Valido(t *testing.T) {
	conectionString := "user=postgres dbname=personas_db_testing host=localhost sslmode=disable password=kys0128tomas port=5432"
	db, err := ConectionDBTesting(conectionString)
	if err != nil {
		t.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Obtener las personas de la base de datos
	personas, err := db_.ObtenerPersonas(db)
	if err != nil {
		t.Fatalf("Error al obtener las personas de la base de datos: %v", err)
	}

	// Verificar que se obtuvieron algunas personas
	if len(personas) == 0 {
		t.Fatalf("No se encontraron personas en la base de datos")
	}
}

func TestObtenerPersonas_Fallido(t *testing.T) {
}

func TestModificarPersonaDB_Valido(t *testing.T) {
	conectionString := "user=postgres dbname=personas_db_testing host=localhost sslmode=disable password=kys0128tomas port=5432"
	db, err := ConectionDBTesting(conectionString)
	if err != nil {
		t.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Persona existente para modificar
	personaInicial := models.Persona{Nombre: "NombreInicial", Apellido: "ApellidoInicial", Edad: 30, CountryCode: "US"}

	//Accion que se encuentra en persona.service.go
	if !personaInicial.Validate() {
		t.Fatalf("Error: %v", &models.ErrorPersonaInvalida{Mensaje: "La persona no es válida"})
	}

	id, err := db_.CreatePersona(db, personaInicial)
	if err != nil {
		t.Fatalf("Error al crear la persona en la base de datos: %v", err)
	}

	//Accion que se encuentra en persona.service.go
	personaInicial.ID = id

	// Obtener la persona recién agregada por su ID
	_, err = db_.ObtenerPersonaDB(db, id)
	if err != nil {
		t.Fatalf("Error al obtener la persona por ID: %v", err)
	}

	// Persona con los nuevos datos
	personaNueva := models.Persona{ID: id, Nombre: "NuevoNombre", Apellido: "NuevoApellido", Edad: 25, CountryCode: "UK"}
	// Modificar persona en la base de datos

	personaModificada, err := ModificarPersonaForTesting(db, personaNueva, personaInicial)
	if err != nil {
		t.Fatalf("Error al modificar persona en la base de datos: %v", err)
	}

	// Verificar que la persona modificada tenga los nuevos datos
	if personaModificada.Nombre != personaNueva.Nombre || personaModificada.Apellido != personaNueva.Apellido || personaModificada.Edad != personaNueva.Edad || personaModificada.CountryCode != personaNueva.CountryCode {
		t.Errorf("Persona modificada incorrectamente. Se esperaba %v pero se obtuvo %v", personaNueva, personaModificada)
	}
}
