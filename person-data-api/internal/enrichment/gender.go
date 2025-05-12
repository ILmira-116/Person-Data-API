package enrichment

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Структура для парсинга ответа от API
type GenderResponse struct {
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

// Функция для получения пола по имени
func GetGender(name string) (string, error) {
	url := fmt.Sprintf("https://api.genderize.io/?name=%s", name)
	log.Infof("Sending request to Genderize API for name: %s", name)

	resp, err := http.Get(url)
	if err != nil {
		log.Debugf("Failed to make request to Genderize API: %v", err)
		return "", fmt.Errorf("failed to make request to genderize: %w", err)
	}
	defer resp.Body.Close()

	var genderResponse GenderResponse
	if err := json.NewDecoder(resp.Body).Decode(&genderResponse); err != nil {
		log.Debugf("Failed to decode Genderize API response: %v", err)
		return "", fmt.Errorf("failed to decode genderize response: %w", err)
	}

	log.Infof("Received gender from Genderize API: name=%s, gender=%s, probability=%.2f",
		genderResponse.Name, genderResponse.Gender, genderResponse.Probability)

	return genderResponse.Gender, nil
}
