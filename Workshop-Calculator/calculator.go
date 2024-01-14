package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(prompt string) float64 {
	fmt.Printf("%v", prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message, _ := fmt.Scanf("%v must number only", prompt)
		panic(message)
	}

	return value

}

func getOperator() string {
	fmt.Print("operator is (+ - * /):")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}

func minus(value1, value2 float64) float64 {
	return value1 - value2
}

func multiply(value1, value2 float64) float64 {
	return value1 * value2
}

func divide(value1, value2 float64) float64 {
	return value1 / value2
}
func main() {
	value1 := getInput("value1 = ")
	value2 := getInput("value2 = ")

	var result float64

	switch operator := getOperator(); operator {
	case "+":
		result = add(value1, value2)
	case "-":
		result = minus(value1, value2)
	case "*":
		result = multiply(value1, value2)
	case "/":
		result = divide(value1, value2)
	default:
		panic("wrong operator")
	}
	fmt.Printf("result is %v", result)
}
