package main

import "fmt"

func add(value1, value2 float64) {
	result := value1 + value2
	fmt.Println("result: ", result)
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}
}

func deferloop() {
	for j := 0; j < 10; j++ {
		defer fmt.Println("j = ", j)
	}
}
func main() {
	// defer fmt.Println("End")
	// fmt.Println("Welcome to Calculator")
	// defer add(10, 20)
	// defer add(63, 43)
	// defer add(94, 58)

	loop()
	deferloop()
}
