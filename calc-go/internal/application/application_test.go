package application_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gromova-aan/Golang/calc-go/internal/application"
	"github.com/gromova-aan/Golang/calc-go/response"
	"github.com/stretchr/testify/assert"
)

func TestCalculateHandler(t *testing.T) {
	// Создаем конфигурацию и приложение
	config := &application.Config{Addr: "8080"}
	app := application.New(config)

	// Создаем валидный запрос
	validRequest := response.RequestBody{
		Expression: "3 + 5",
	}

	reqBody, err := json.Marshal(validRequest)
	if err != nil {
		t.Fatalf("Error marshalling request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(reqBody))
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(app.CalculateHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var respBody response.ResponseBody
	err = json.NewDecoder(rr.Body).Decode(&respBody)
	if err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}

	expectedResult := "8.00"
	assert.Equal(t, expectedResult, respBody.Result)
}

func TestInvalidExpression(t *testing.T) {
	// Создаем конфигурацию и приложение
	config := &application.Config{Addr: "8080"}
	app := application.New(config)

	// Создаем некорректный запрос
	invalidRequest := response.RequestBody{
		Expression: "10 / 0",
	}

	reqBody, err := json.Marshal(invalidRequest)
	if err != nil {
		t.Fatalf("Error marshalling request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(reqBody))
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(app.CalculateHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)

	var respBody response.ResponseBody
	err = json.NewDecoder(rr.Body).Decode(&respBody)
	if err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}

	expectedError := "division by zero"
	assert.Equal(t, expectedError, respBody.Error)
}

func TestInvalidMethod(t *testing.T) {
	// Создаем конфигурацию и приложение
	config := &application.Config{Addr: "8080"}
	app := application.New(config)

	// Отправляем запрос с некорректным методом
	req := httptest.NewRequest(http.MethodGet, "/api/v1/calculate", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(app.CalculateHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
}

func TestEmptyExpression(t *testing.T) {
	// Создаем конфигурацию и приложение
	config := &application.Config{Addr: "8080"}
	app := application.New(config)

	// Создаем запрос с пустым выражением
	emptyRequest := response.RequestBody{
		Expression: "",
	}

	reqBody, err := json.Marshal(emptyRequest)
	if err != nil {
		t.Fatalf("Error marshalling request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(reqBody))
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(app.CalculateHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)

	var respBody response.ResponseBody
	err = json.NewDecoder(rr.Body).Decode(&respBody)
	if err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}

	expectedError := "Expression is not valid"
	assert.Equal(t, expectedError, respBody.Error)
}
