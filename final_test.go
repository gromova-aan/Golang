package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var priority = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
	"(": 0,
}

func isOperator(token string) bool {
	_, exists := priority[token]
	return exists
}

func ToPostfix(expression string) (string, error) {
	// Регулярное выражение для поиска чисел и операторов, включая скобки
	re := regexp.MustCompile(`\d+(\.\d+)?|[+\-*/()]`)
	tokens := re.FindAllString(expression, -1)

	var output []string
	var operators []string

	for _, token := range tokens {
		if _, err := strconv.ParseFloat(token, 64); err == nil {
			output = append(output, token)
		} else if token == "(" {
			operators = append(operators, token)
		} else if token == ")" {
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			if len(operators) == 0 {
				return "", errors.New("несоответствующая скобка")
			}
			operators = operators[:len(operators)-1] // Удаляем "("
		} else if isOperator(token) {
			for len(operators) > 0 && priority[operators[len(operators)-1]] >= priority[token] {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		} else {
			return "", errors.New("недопустимый символ в выражении")
		}
	}

	for len(operators) > 0 {
		if operators[len(operators)-1] == "(" {
			return "", errors.New("несоответствующая скобка")
		}
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	return strings.Join(output, " "), nil
}

func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")
	
	// Преобразуем инфиксное выражение в постфиксное
	postfix, err := ToPostfix(expression)
	if err != nil {
		return 0, err
	}

	tokens := strings.Fields(postfix)
	var stack []float64

	for _, token := range tokens {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, num)
		} else {
			// Проверяем, что в стеке есть хотя бы два числа
			if len(stack) < 2 {
				return 0, fmt.Errorf("недостаточно операндов для операции %s", token)
			}

			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result float64
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("деление на ноль")
				}
				result = a / b
			default:
				return 0, fmt.Errorf("неизвестный оператор: %s", token)
			}

			// Добавляем результат обратно в стек
			stack = append(stack, result)
		}
	}

	// В стеке должен остаться один элемент — результат
	if len(stack) != 1 {
		return 0, fmt.Errorf("ошибка в выражении")
	}

	return stack[0], nil
}

func main() {
	expression := "3 + 4 * 2.2 / ( 5.5 - 5 )"
	result, err := Calc(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Результат: %.2f\n", result)
	}
}