package db

import (
	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

// DeletePersonDb удаляет человека из базы данных по его ID
func DeletePersonDb(id int) error {
	log.Infof("Attempting to delete person with ID: %d", id)

	// Формируем SQL-запрос на удаление
	query := "DELETE FROM people WHERE id = $1"
	log.Debugf("Executing delete query: %s with ID: %d", query, id)

	// Выполняем запрос
	result, err := DB.Exec(query, id)
	if err != nil {
		log.Infof("Error executing delete query: %v", err)
		return err
	}

	// Проверяем, сколько строк было затронуто
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Infof("Error getting rows affected: %v", err)
		return err
	}

	// Логируем результат
	if rowsAffected == 0 {
		log.Infof("No person found with ID: %d", id)
	} else {
		log.Infof("Successfully deleted person with ID: %d (Rows affected: %d)", id, rowsAffected)
	}

	return nil
}
