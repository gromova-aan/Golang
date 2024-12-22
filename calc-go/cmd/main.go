package main

import (
	"log"

	"github.com/gromova-aan/Golang/calc-go/internal/application"
)

func main() {
	// Инициализация конфигурации
	config := application.ConfigFromEnv()

	// Создаем приложение
	app := application.New(config)

	// Запускаем приложение
	if err := app.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
