package application

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	rpn "github.com/gromova-aan/Golang/calc-go/pkg/calculation"
	"github.com/gromova-aan/Golang/calc-go/response"
)

// Config для хранения конфигурации приложения
type Config struct {
	Addr string
}

// ConfigFromEnv загружает конфигурацию из переменных окружения
func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

// Application представляет приложение
type Application struct {
	Config *Config
}

// New создает новый экземпляр Application
func New(config *Config) *Application {
	return &Application{Config: config}
}

// Run запускает сервер приложения
func (a *Application) Run() error {
	log.Printf("Server is running on port %s", a.Config.Addr)
	return http.ListenAndServe(":"+a.Config.Addr, a.routes())
}

// routes задает маршруты для приложения
func (a *Application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", a.CalculateHandler)
	return mux
}

// CalculateHandler обработчик запроса на вычисление
func (a *Application) CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody response.RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		a.sendJSONResponse(w, http.StatusInternalServerError, response.ResponseBody{Error: "Internal server error"})
		return
	}

	if reqBody.Expression == "" {
		a.sendJSONResponse(w, http.StatusUnprocessableEntity, response.ResponseBody{Error: "Expression is not valid"})
		return
	}

	result, calcErr := rpn.Calc(reqBody.Expression)
	if calcErr != nil {
		a.sendJSONResponse(w, http.StatusUnprocessableEntity, response.ResponseBody{Error: calcErr.Error()})
		return
	}

	a.sendJSONResponse(w, http.StatusOK, response.ResponseBody{Result: strconv.FormatFloat(result, 'f', 2, 64)})
}

// sendJSONResponse отправляет JSON-ответ
func (a *Application) sendJSONResponse(w http.ResponseWriter, statusCode int, respBody response.ResponseBody) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(respBody)
}
