package db

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv" // подключаем godotenv для работы с .env
	_ "github.com/lib/pq"      // подключение драйвера для Postgres
)

// DB переменная для глобального доступа к базе данных
var DB *sqlx.DB

func InitDB() {
	// Загружаем переменные окружения из .env файла
	log.Debug("Loading environment variables from .env file")
	err := godotenv.Load()
	if err != nil {
		log.Warn("Error loading .env file")
	}

	// Получаем DSN для подключения к базе данных из переменной окружения
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Warn("Database connection string is empty")
		return
	}

	// Логируем информацию о подключении к базе данных
	log.Info("Connecting to the database...")

	// Подключение к базе данных Postgres с использованием DSN
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Warnf("Failed to connect to the database: %v", err)
		return
	}

	// Логируем успешное подключение
	log.Info("Successfully connected to the database")
}
