package db

import (
	"fmt"
	"person-data-api/internal/model"

	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

func UpdatePersonDb(id int, p model.Person) error {
	log.Infof("Attempting to update person with ID: %d", id)

	// Формируем SQL-запрос на удаление
	query := `UPDATE people 
	SET name = $1,
		surname = $2,
		patronymic = $3,
		updated_at = NOW()
	WHERE id = $4;`

	// Выполняем запрос
	result, err := DB.Exec(query, p.Name, p.Surname, p.Patronymic, id)
	if err != nil {
		log.Debugf("Update query execution failed for ID %d: %v", id, err)
		return err
	}

	// Проверяем, сколько строк было затронуто
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Debugf("Failed to retrieve rows affected for ID %d: %v", id, err)
		return err
	}

	// Логируем результат
	if rowsAffected == 0 {
		log.Infof("No person found with ID: %d", id)
		return fmt.Errorf("person with ID %d not found", id)
	} else {
		log.Infof("Successfully updated person with ID: %d (Rows affected: %d)", id, rowsAffected)
	}

	return nil
}
