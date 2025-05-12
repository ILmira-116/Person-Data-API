package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Инициализация логгера с нужными настройками
func Init() {
	// Устанавливаем форматтер для логов
	log.SetFormatter(&log.TextFormatter{
		// Включаем полное отображение времени в логах
		FullTimestamp: true,
	})

	// Устанавливаем вывод логов в стандартный вывод (консоль)
	log.SetOutput(os.Stdout)

	// Устанавливаем уровень логирования.
	log.SetLevel(log.DebugLevel) // можно менять: InfoLevel, ErrorLevel и т.д.
}
