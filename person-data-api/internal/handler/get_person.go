package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"person-data-api/internal/db"
	"strconv"

	_ "person-data-api/docs"

	log "github.com/sirupsen/logrus"
)

// GetPerson godoc
// @Summary Получение списка людей
// @Description Возвращает список людей с пагинацией
// @Tags people
// @Param page query int false "Номер страницы (по умолчанию 1)"
// @Param limit query int false "Количество элементов на странице (по умолчанию 10)"
// @Success 200 {object} model.Person
// @Failure 400 {string} string "Некорректные параметры запроса"
// @Failure 500 {string} string "Ошибка при получении данных из базы"
// @Router /people [get]
func GetPerson(w http.ResponseWriter, r *http.Request) {
	log.Info("Received request to get persons")

	// Извлекаем параметры пагинации из URL
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	if page == "" {
		page = "1" // По умолчанию, если не указан page, то считаем, что это 1
	}

	if limit == "" {
		limit = "10" // По умолчанию, если не указан limit, то считаем, что на странице 10 записей
	}

	log.Debugf("Pagination parameters: page = %s, limit = %s", page, limit)

	pageNum, err := strconv.Atoi(page) // Преобразуем номер страницы в int
	if err != nil {
		log.Debugf("Invalid page number: %v", err)
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	limitNum, err := strconv.Atoi(limit) // Преобразуем лимит в int
	if err != nil {
		log.Debugf("Invalid limit: %v", err)
		http.Error(w, "Invalid limit", http.StatusBadRequest)
		return
	}

	// Расчет OFFSET
	offset := (pageNum - 1) * limitNum
	log.Debugf("Calculated offset: %d", offset)

	// Получаем список пользователей из базы данных
	persons, err := db.GetAllPerson(offset, limitNum)
	if err != nil {
		log.Debugf("Error fetching data from database: %v", err)
		http.Error(w, fmt.Sprintf("Error fetching data: %v", err), http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок для ответа
	w.Header().Set("Content-Type", "application/json")

	// Сериализуем список пользователей в JSON
	response, err := json.Marshal(persons)
	if err != nil {
		log.Debugf("Error marshaling data: %v", err)
		http.Error(w, fmt.Sprintf("Error marshaling data: %v", err), http.StatusInternalServerError)
		return
	}

	// Отправляем ответ
	log.Infof("Successfully fetched %d persons", len(persons))
	w.Write(response)
}
