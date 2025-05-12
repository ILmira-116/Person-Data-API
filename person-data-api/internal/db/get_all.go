package db

import (
	"person-data-api/internal/model"

	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

// GetAllPerson получает список людей из базы данных с учётом пагинации
func GetAllPerson(offset, limitNum int) ([]model.Person, error) {
	log.Debugf("Fetching people from DB with limit: %d and offset: %d", limitNum, offset)

	// Формируем запрос в базу данных
	query := "SELECT id, name, surname, patronymic, age, gender, nationality, created_at, updated_at FROM people LIMIT $1 OFFSET $2"
	log.Debug("Executing SQL query to fetch people")

	// Получаем строки из базы данных с параметризованным запросом
	rows, err := DB.Query(query, limitNum, offset)
	if err != nil {
		log.Errorf("Failed to execute query: %v", err)
		return nil, err
	}
	defer rows.Close()
	log.Debug("Query executed successfully, processing rows")

	// Проходимся по всем стркоам
	var persons []model.Person
	for rows.Next() {
		var p model.Person
		if err := rows.Scan(
			&p.ID, &p.Name, &p.Surname, &p.Patronymic,
			&p.Age, &p.Gender, &p.Nationality, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			log.Errorf("Error scanning row: %v", err)
			return nil, err
		}
		persons = append(persons, p)
	}

	if err := rows.Err(); err != nil {
		log.Infof("Rows iteration finished with error: %v", err)
	}

	log.Infof("Successfully fetched %d people from DB", len(persons))
	return persons, nil
}
