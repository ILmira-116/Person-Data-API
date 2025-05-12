package handler

import (
	"encoding/json"
	"net/http"
	"person-data-api/internal/db"
	"person-data-api/internal/model"
	"strconv"
	"strings"

	_ "person-data-api/docs"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// UpdatePerson godoc
// @Summary Обновление данных о человеке (частичное обновление)
// @Description Частично обновляет данные о человеке по его ID в базе данных
// @Tags people
// @Param id path int true "ID человека для обновления"
// @Param person body model.Person true "Данные человека для частичного обновления"
// @Success 200 {string} string "Person updated successfully"
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Person with this ID not found"
// @Failure 500 {string} string "Failed to update person"
// @Router /people/{id} [patch]
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	log.Info("Received request to update a person")

	// Получаем ID из параметров URL
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		log.Debug("Missing 'id' parameter in URL path")
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	// Преобразуем строку в int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Debugf("Invalid ID format: %v", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	log.Debugf("Updating person with ID: %d", id)

	//Читаем JSON из тела запроса
	var p model.Person
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Debugf("Invalid JSON format: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Debugf("Decoded person data: %+v", p)

	// Вызываем функцию для обновления данных в БД
	err = db.UpdatePersonDb(id, p)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, "Person with this ID not found", http.StatusNotFound)
			return
		}
		log.Debugf("Failed to update person with ID %d: %v", id, err)
		http.Error(w, "Failed to update person", http.StatusInternalServerError)
		return
	}

	log.Infof("Successfully updated person with ID %d", id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Person updated successfully"))
}
