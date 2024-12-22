package rpn_test

import (
	"testing"

	rpn "github.com/gromova-aan/Golang/calc-go/pkg/calculation"
)

func TestCalc(t *testing.T) {
	// Тестовые случаи для успешных вычислений
	testCasesSuccess := []struct {
		name           string
		expression     string
		expectedResult float64
	}{
		{
			name:           "Simple addition",
			expression:     "1+1",
			expectedResult: 2,
		},
		{
			name:           "Priority with parentheses",
			expression:     "(2+2)*2",
			expectedResult: 8,
		},
		{
			name:           "Operator precedence",
			expression:     "2+2*2",
			expectedResult: 6,
		},
		{
			name:           "Division",
			expression:     "1/2",
			expectedResult: 0.5,
		},
		{
			name:           "Complex expression",
			expression:     "3+(4*2)-(10/5)",
			expectedResult: 8,
		},
	}

	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := rpn.Calc(testCase.expression)
			if err != nil {
				t.Fatalf("Expected no error for case %s, but got: %v", testCase.expression, err)
			}
			if val != testCase.expectedResult {
				t.Fatalf("Expected %f for case %s, but got %f", testCase.expectedResult, testCase.expression, val)
			}
		})
	}

	// Тестовые случаи для некорректных выражений
	testCasesFail := []struct {
		name        string
		expression  string
		expectedErr string
	}{
		{
			name:       "Invalid operator placement",
			expression: "1+1*",
			expectedErr: "error in expression",
		},
		{
			name:       "Double operators",
			expression: "2+2**2",
			expectedErr: "error in expression",
		},
		{
			name:       "Unmatched parentheses",
			expression: "((2+2)-(2",
			expectedErr: "mismatched bracket",
		},
		{
			name:       "Empty expression",
			expression: "",
			expectedErr: "invalid character in expression",
		},
		{
			name:       "Division by zero",
			expression: "10/0",
			expectedErr: "division by zero",
		},
	}

	for _, testCase := range testCasesFail {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := rpn.Calc(testCase.expression)
			if err == nil {
				t.Fatalf("Expected error for expression %s, but got no error", testCase.expression)
			}
			if err.Error() != testCase.expectedErr {
				t.Fatalf("Expected error %s for expression %s, but got %s", testCase.expectedErr, testCase.expression, err.Error())
			}
		})
	}
}
