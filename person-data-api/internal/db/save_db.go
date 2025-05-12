package db

import (
	"fmt"
	"person-data-api/internal/model"

	log "github.com/sirupsen/logrus"
)

// SavePerson сохраняет информацию о человеке в базе данных
func SavePerson(person *model.Person) error {
	// Логируем параметры перед вставкой
	log.Debugf("Attempting to insert person into DB: %+v", person)

	// SQL-запрос для вставки данных в таблицу persons
	query :=
		`INSERT INTO people(name, surname, patronymic, age, gender, nationality, created_at, updated_at)
	VALUES (:name, :surname, :patronymic, :age, :gender, :nationality, :created_at, :updated_at)`

	// Логируем начало операции
	log.Debugf("Preparing to insert person: %+v", person)

	// Выполняем запрос через NamedQuery, передаём сразу всю структуру `person`.
	_, err := DB.NamedExec(query, person)
	if err != nil {
		log.Debugf("Insert failed for person: %+v, error: %v", person, err)
		return fmt.Errorf("could not insert person: %w", err)
	}

	// Логируем успешную вставку
	log.Infof("Successfully inserted person: %s %s", person.Name, person.Surname)
	return nil
}
