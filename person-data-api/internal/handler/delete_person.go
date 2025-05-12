package handler

import (
	"net/http"
	"person-data-api/internal/db"
	"strconv"

	_ "person-data-api/docs"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// DeletePerson godoc
// @Summary Удаление человека
// @Description Удаляет человека по ID
// @Tags people
// @Param id path int true "ID человека"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Некорректный ID"
// @Failure 500 {string} string "Ошибка удаления из базы данных"
// @Router /people/{id} [delete]
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	log.Info("Received request to delete a person")

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
		log.Infof("Invalid ID format: %s", idStr)
		log.Debugf("Conversion error: %v", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	//Удаляем пользователя из базы данных
	err = db.DeletePersonDb(id)
	if err != nil {
		log.Infof("Failed to delete person with ID %d", id)
		log.Debugf("DB deletion error: %v", err)
		http.Error(w, "Error deleting person", http.StatusInternalServerError)
		return
	}

	log.Infof("Successfully deleted person with ID %d", id)
	w.WriteHeader(http.StatusNoContent) // 204 No Content — успешно, но тело не возвращается
}
