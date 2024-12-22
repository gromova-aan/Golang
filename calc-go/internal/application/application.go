package application

import (
	"encoding/json"
	"net/http"
	"strconv"

	rpn "github.com/gromova-aan/Golang/calc-go/pkg/calculation"
	"github.com/gromova-aan/Golang/calc-go/response"
)

// Обработчик запроса на вычисление
func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что запрос POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Читаем тело запроса
	var reqBody response.RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		sendJSONResponse(w, http.StatusInternalServerError, response.ResponseBody{Error: "Internal server error"}) //422
		return
	}

	// Проверяем, что выражение не пустое
	if reqBody.Expression == "" {
		sendJSONResponse(w, http.StatusUnprocessableEntity, response.ResponseBody{Error: "Expression is not valid"}) //200
		return
	}

	// Проверяем на допустимые символы и вычисляем выражение
	result, calcErr := rpn.Calc(reqBody.Expression)
	if calcErr != nil {
		sendJSONResponse(w, http.StatusUnprocessableEntity, response.ResponseBody{Error: calcErr.Error()})
		return
	}

	// Возвращаем успешный ответ
	sendJSONResponse(w, http.StatusOK, response.ResponseBody{Result: strconv.FormatFloat(result, 'f', 2, 64)})
}

// Общая функция для отправки JSON-ответа
func sendJSONResponse(w http.ResponseWriter, statusCode int, respBody response.ResponseBody) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(respBody)
}
