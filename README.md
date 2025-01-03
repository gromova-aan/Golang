# Калькулятор

Это простой API калькулятор, разработанный на языке Go, который позволяет пользователям выполнять базовые арифметические операции, такие как сложение, вычитание, умножение и деление. Он поддерживает обратную польскую запись (ОПЗ) для парсинга и вычисления выражений.

## Возможности

- Поддержка базовых арифметических операций: `+`, `-`, `*`, `/`
- Обработка выражений в обратной польской записи (ОПЗ)
- Обработка ошибок (например, деление на ноль, недопустимые символы)
- RESTful API, построенный с использованием пакета `net/http` в Go
- Юнит-тесты с использованием тестового фреймворка Go

## Требования

Перед запуском проекта убедитесь, что у вас установлены:

- Go (версия 1.17 и выше)
- Git

## Установка

1. Клонируйте этот репозиторий:

    ```bash
    git clone https://github.com/gromova-aan/Golang.git
    ```

2. Перейдите в каталог проекта:

    ```bash
    cd calc-go
    ```

3. Инициализируйте Go-модуль (если этого еще не сделано):

    ```bash
    go mod init calc-go
    ```

4. Установите зависимости:

    ```bash
    go get github.com/stretchr/testify/assert
    ```


## Запуск проекта

Для запуска API-сервера выполните следующую команду:

```bash
go run cmd/main.go
```
## API Эндпоинты
POST /api/v1/calculate

Этот эндпоинт принимает POST-запрос с JSON-телом, содержащим математическое выражение в инфиксной нотации. Сервер вычисляет выражение и возвращает результат в формате JSON.

Пример тела запроса:

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```
Пример тела ответа: код 200 ОК

```bash
{
  "result": "6.00"
}
```
Ошибки:

400 Bad Request: Если выражение пустое или содержит недопустимые символы.

422 Unprocessable Entity: Если выражение некорректное или вызывает ошибку (например, деление на ноль).
Пример ошибки:

```bash
{
  "error": "division by zero"
}
```
500 в случае какой-либо иной ошибки («Что-то пошло не так»).
```bash
{
    "error": "Internal server error"
}
```

Запуск тестов
Для запуска тестов проекта выполните следующую команду:

```bash
 cd calc-go/internal/application/
go test -v
```
Эта команда выполнит тесты и выведет подробный результат.

Вы также можете запустить конкретные тесты:

```bash
go test -v -run TestCalculateHandler
```

Структура проекта:
```bash
calc-go/
│
├── cmd/
│   └── main.go                # Точка входа в приложение (запуск сервера)
│
├── internal/
│   └── application/
│       └── application.go     # Логика обработки запросов (обработчики)
│
├── pkg/
│   └── calculation/
│       ├── calculation.go             # Логика вычислений
│       ├── calculation_test.go        # Тесты для вычислений
│
└── response/
    └── response.go            # Структуры запросов и ответов
```
