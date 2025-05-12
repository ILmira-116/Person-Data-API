package handler

import (
	"encoding/json"
	"fmt"

	"net/http"
	"person-data-api/internal/client"
	"person-data-api/internal/db"
	"person-data-api/internal/payload"

	_ "person-data-api/docs"

	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

var validate = validator.New()

// AddPerson godoc
// @Summary Добавить нового человека
// @Description Принимает данные о человеке, обогащает их через внешний API и сохраняет в базу данных
// @Tags person
// @Accept json
// @Produce json
// @Param person body payload.PersonPayload true "Данные человека"
// @Success 201 {object} map[string]string "Person successfully added"
// @Failure 400 {string} string "Invalid JSON / Validation error / Error enriching person data"
// @Failure 500 {string} string "Failed to save person to DB"
// @Router /person [post]
func AddPerson(w http.ResponseWriter, r *http.Request) {
	// Логируем информацию о полученном запросе
	log.Info("Received a POST request to add a person")

	// Читаем тело запроса (payload)
	var person payload.PersonPayload
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		log.Debugf("Invalid JSON body: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Проверяем данные на соответствие структуре
	err = validate.Struct(person)
	if err != nil {
		log.Debugf("Validation error: %v", err)
		http.Error(w, fmt.Sprintf("Validation error: %v", err), http.StatusBadRequest)
		return
	}

	log.Debugf("Received payload: %+v", person)

	// Выводим данные, которые получили в payload
	fmt.Printf("Received Person: Name: %s, Surname: %s, Patronymic: %s\n", person.Name, person.Surname, person.Patronymic)

	// Обогащаем клиента данными
	enrichedPerson, err := client.GetPersonData(person.Name, person.Surname, person.Patronymic)
	if err != nil {
		log.Infof("Error enriching person data for name: %s", person.Name)
		log.Debugf("GetPersonData error: %v", err)
		http.Error(w, "Error enriching person data", http.StatusBadRequest)
		return
	}
	log.Infof("Successfully enriched person: %s %s", person.Name, person.Surname)
	log.Debugf("Enriched person details: %+v", enrichedPerson)

	// Сохраняем обогащенные данные в БД
	if err := db.SavePerson(enrichedPerson); err != nil {
		log.Infof("Failed to save person to DB")
		log.Debugf("DB error: %v", err)
		http.Error(w, "Failed to save person to DB", http.StatusInternalServerError)
		return
	}

	log.Info("Person successfully saved to DB")

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Формируем ответ в формате JSON
	response := map[string]string{
		"message": "Person successfully added",
	}

	// Отправляем JSON-ответ с кодом 201 (создано)
	w.WriteHeader(http.StatusCreated)
	jsonResponse, _ := json.Marshal(response)
	w.Write(jsonResponse)
}
