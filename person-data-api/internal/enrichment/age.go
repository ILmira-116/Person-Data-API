package enrichment

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Структура для парсинга ответа от API
type AgeResponse struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Count int    `json:"count"`
}

// Функция для получения возраста по имени
func GetAge(name string) (int, error) {
	// формируем url по которому будем обращаться к api
	url := fmt.Sprintf("https://api.agify.io/?name=%s", name)
	log.Infof("Sending request to Agify API for name: %s", name)

	// Выполняем GET-запрос
	resp, err := http.Get(url)
	if err != nil {
		log.Debugf("Failed to make request to Agify API: %v", err)
		return 0, fmt.Errorf("failed to make request to agify: %w", err)
	}
	defer resp.Body.Close()

	// Создаём переменную для хранения ответа от API
	var ageResponse AgeResponse

	// Декодируем JSON-ответ в структуру AgeResponse
	if err := json.NewDecoder(resp.Body).Decode(&ageResponse); err != nil {
		log.Debugf("Failed to decode Agify API response: %v", err)
		return 0, fmt.Errorf("failed to decode agify response: %w", err)
	}

	// Возвращаем возраст, полученный от API
	log.Infof("Received age from Agify API: name=%s, age=%d", ageResponse.Name, ageResponse.Age)
	return ageResponse.Age, nil
}
