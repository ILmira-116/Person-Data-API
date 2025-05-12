// Package client отвечает за получение и обогащение данных о человеке
package client

import (
	"fmt"
	"person-data-api/internal/enrichment"
	"person-data-api/internal/model"
	"time"

	log "github.com/sirupsen/logrus"
)

// GetPersonData запрашивает и собирает данные о человеке:
// возраст, пол и национальность на основе имени
func GetPersonData(name, surname, patronymic string) (*model.Person, error) {
	log.Infof("Starting data enrichment for person: %s %s %s", surname, name, patronymic)

	var person model.Person

	// Получение возраста
	log.Debugf("Fetching age for name: %s", name)
	age, err := enrichment.GetAge(name)
	if err != nil {
		log.Errorf("Failed to get age: %v", err)
		return nil, fmt.Errorf("error getting age: %w", err)
	}
	log.Debugf("Retrieved age: %d", age)

	// Получение пола
	log.Debugf("Fetching gender for name: %s", name)
	gender, err := enrichment.GetGender(name)
	if err != nil {
		log.Errorf("Failed to get gender: %v", err)
		return nil, fmt.Errorf("error getting gender: %w", err)
	}
	log.Debugf("Retrieved gender: %s", gender)

	// Получение национальности
	log.Debugf("Fetching nationality for name: %s", name)
	nationality, err := enrichment.GetNationality(name)
	if err != nil {
		log.Errorf("Failed to get nationality: %v", err)
		return nil, fmt.Errorf("error getting nationality: %w", err)
	}
	log.Debugf("Retrieved nationality: %s", nationality)

	// Заполнение модели Person
	person = model.Person{
		Name:        name,
		Surname:     surname,
		Patronymic:  patronymic,
		Age:         age,
		Gender:      gender,
		Nationality: nationality,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	log.Infof("Successfully enriched person data: %+v", person)

	return &person, nil
}
