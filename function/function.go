package main

import "fmt"

func hello() {
	fmt.Println("Hello The.Dome")
}

func plus(value1 int, value2 int) int {
	//result := value1 + value2
	//fmt.Println("result = ", result)
	return value1 + value2
}

func plus3value(value1, value2, value3 int) int {
	return value1 + value2 + value3
}
func main() {
	hello()
	//plus(1, 2)
	result := plus(4, 6)
	fmt.Println("result =", result)
	result2 := plus3value(5, 2, 9)
	fmt.Println("result2 = ", result2)
}
