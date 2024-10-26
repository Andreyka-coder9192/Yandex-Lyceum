package main

import (
	"errors"
	"fmt"
	"strconv"
)

func Calc(expression string) (float64, error) {
	// преобразование выражения в обратную польскую нотацию
	rpn, err := infixToPostfix(expression)
	if err != nil {
		return 0, err
	}

	// вычисление результата из обратной польской нотации
	result, err := evaluatePostfix(rpn)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func infixToPostfix(expression string) ([]string, error) {
	var rpn []string
	var stack []string

	precedence := map[string]int{"+": 1, "-": 1, "*": 2, "/": 2}

	tokens := tokenize(expression)

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[token] {
				rpn = append(rpn, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		case "(":
			stack = append(stack, token)
		case ")":
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				rpn = append(rpn, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil, errors.New("mismatched parentheses")
			}
			stack = stack[:len(stack)-1]
		default:
			rpn = append(rpn, token)
		}
	}

	for len(stack) > 0 {
		if stack[len(stack)-1] == "(" {
			return nil, errors.New("mismatched parentheses")
		}
		rpn = append(rpn, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return rpn, nil
}

func evaluatePostfix(rpn []string) (float64, error) {
	var stack []float64

	for _, token := range rpn {
		switch token {
		case "+":
			if len(stack) < 2 {
				return 0, errors.New("not enough operands for + operator")
			}
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, operand1+operand2)
		case "-":
			if len(stack) < 2 {
				return 0, errors.New("not enough operands for - operator")
			}
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, operand1-operand2)
		case "*":
			if len(stack) < 2 {
				return 0, errors.New("not enough operands for * operator")
			}
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, operand1*operand2)
		case "/":
			if len(stack) < 2 {
				return 0, errors.New("not enough operands for / operator")
			}
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			if operand2 == 0 {
				return 0, errors.New("division by zero")
			}
			stack = stack[:len(stack)-2]
			stack = append(stack, operand1/operand2)
		default:
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, errors.New("invalid operand")
			}
			stack = append(stack, num)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}

	return stack[0], nil
}

func tokenize(expression string) []string {
	var tokens []string
	num := ""
	for _, char := range expression {
		if char >= '0' && char <= '9' || char == '.' {
			num += string(char)
		} else {
			if num != "" {
				tokens = append(tokens, num)
				num = ""
			}
			if char == '(' || char == ')' || char == '+' || char == '-' || char == '*' || char == '/' {
				tokens = append(tokens, string(char))
			}
		}
	}
	if num != "" {
		tokens = append(tokens, num)
	}
	return tokens
}

func main() {
	result, err := Calc("1/2*4+8")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

// for identif != -1 {
// 	for pos, val := range expression {
// 		if priority(string(val)) != identif {
// 			continue
// 		}

// 		if priority(string(val)) == 2 {
// 			for pos1, val1 := range expression {
// 				if priority(string(val1)) == 2 {
// 					skobki = append(skobki, pos)
// 					skobki = append(skobki, pos1)
// 				}
// 			}
// 		}
// 	}
// 	identif--
// }
