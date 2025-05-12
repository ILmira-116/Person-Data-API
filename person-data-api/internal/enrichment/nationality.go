package enrichment

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Структура для парсинга ответа от API
type NationalityResponse struct {
	Name    string               `json:"name"`
	Country []CountryProbability `json:"country"`
}

// Структура обработки  одного элемента в массиве
type CountryProbability struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

// Функция для получения национальности по имени
func GetNationality(name string) (string, error) {
	// формируем url по которому будем обращаться к api
	url := fmt.Sprintf("https://api.nationalize.io/?name=%s", name)
	log.Infof("Sending request to Nationalize API for name: %s", name)

	// Отправляем запрос Get
	resp, err := http.Get(url)
	if err != nil {
		log.Debugf("Failed to make request to Nationalize API: %v", err)
		return "", fmt.Errorf("failed to make request to nationalize: %w", err)
	}
	defer resp.Body.Close()

	var nationalityResponse NationalityResponse
	if err := json.NewDecoder(resp.Body).Decode(&nationalityResponse); err != nil {
		log.Debugf("Failed to decode Nationalize API response: %v", err)
		return "", fmt.Errorf("failed to decode nationalize response: %w", err)
	}

	if len(nationalityResponse.Country) == 0 {
		log.Infof("No nationality data found for name: %s", name)
		return "unknown", nil
	}

	maxProbality := -1.0
	var maxCountryProbalitity string
	for _, country := range nationalityResponse.Country {
		log.Debugf("Evaluating country: %s with probability %.2f", country.CountryID, country.Probability)
		if country.Probability > maxProbality {
			maxProbality = country.Probability
			maxCountryProbalitity = country.CountryID
		}
	}

	log.Infof("Most probable nationality for name %s is %s with probability %.2f", name, maxCountryProbalitity, maxProbality)
	return maxCountryProbalitity, nil
}
