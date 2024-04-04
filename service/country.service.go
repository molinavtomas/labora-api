package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/molinavtomas/labora-api-personas/models"
)

func getCountryInfo(countryCode string) (models.CountryInfo, error) {
	url := fmt.Sprintf("https://restcountries.com/v3.1/alpha/%s", countryCode)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al hacer la solicitud")
		return models.CountryInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return models.CountryInfo{}, fmt.Errorf("error en la solicitud: %s", resp.Status)
	}

	var countryResponse models.CountryResponse
	err = json.NewDecoder(resp.Body).Decode(&countryResponse)
	if err != nil {
		fmt.Println("Error al decodificar la respuesta")
		return models.CountryInfo{}, err
	}

	// Tomar solo el primer país (suponiendo que solo haya uno en la respuesta)
	country := countryResponse[0]

	//obtener key del mapa

	// Extraer la información necesaria

	countryInfo := models.CountryInfo{
		Name:     country.Name.Common,
		Timezone: country.Timezones[0], // Tomar solo el primer timezone
		Flag:     country.Flags.PNG,
	}

	for key := range country.Currencies {
		countryInfo.Currency = key //obtenemos el simbolo de la moneda
	}

	return countryInfo, nil
}
