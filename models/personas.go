package models

import "fmt"

type Persona struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Apellido    string `json:"apellido"`
	Edad        int    `json:"edad"`
	CountryCode string `json:"countryCode"`
}

type PersonaExtendida struct {
	Persona
	CountryInfo
}

type ErrorPersonaInvalida struct {
	Mensaje string
}

func (e *ErrorPersonaInvalida) Error() string {
	return fmt.Sprintf("Error: Persona inválida. %s", e.Mensaje)
}

func (p *Persona) Validate() bool {
	if p.Nombre == "" || p.Apellido == "" || p.Edad == 0 || p.CountryCode == "" {
		return false
	}
	return true
}
