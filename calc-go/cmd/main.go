package main

import (
	"log"
	"net/http"

	"github.com/gromova-aan/Golang/calc-go/internal/application"
)

func main() {
	// Регистрация обработчика для вычислений
	http.HandleFunc("/api/v1/calculate", application.CalculateHandler)

	// Запуск сервера
	port := ":8080"
	log.Printf("Server is running on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}