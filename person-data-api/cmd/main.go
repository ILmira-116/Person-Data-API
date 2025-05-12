package main

import (
	"net/http"
	"person-data-api/internal/db"
	"person-data-api/internal/handler"
	"person-data-api/internal/logger"

	_ "person-data-api/docs"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Person Data API
// @version 1.0
// @description API для добавления и обработки данных о людях.
// @host localhost:8080
// @BasePath /
func main() {
	// Инициализация логгера
	logger.Init()
	log.Info("Logger initialized")

	// Загрузка .env
	if err := godotenv.Load(); err != nil {
		log.Debug(".env file not found, using default environment variables")
	} else {
		log.Info(".env file loaded successfully")
	}

	// Инициализация подключения к базе данных
	log.Debug("Initializing database connection")
	db.InitDB()
	log.Info("Database initialized")

	// Инициалиация роутера
	router := mux.NewRouter()
	log.Debug("Router initialized")

	// Определяем путь и метод для обработки запроса
	router.HandleFunc("/person", handler.AddPerson).Methods("POST")           // Для добавления нового человека
	router.HandleFunc("/person", handler.GetPerson).Methods("GET")            // Для получения списка людей с фильтрацией и пагинацией
	router.HandleFunc("/person/{id}", handler.DeletePerson).Methods("DELETE") // Для удаления человека по ID
	router.HandleFunc("/person/{id}", handler.UpdatePerson).Methods("PATCH")  // Для изменения данных человека по ID
	log.Info("Routes configured")

	// Статичный путь для Swagger UI
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	// Запуск сервера
	log.Info("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
