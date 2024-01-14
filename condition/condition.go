package main

import "fmt"

var score int

func main() {
	fmt.Println("GRADE CALCULATOR")
	fmt.Scanf("%d", &score)
	fmt.Println("Score = ", score)
	if score >= 80 {
		fmt.Println("GRADE A")
	} else if score >= 70 {
		fmt.Println("GRADE B")
	} else if score >= 60 {
		fmt.Println("GRADE C")
	} else if score >= 50 {
		fmt.Println("GRADE D")
	} else if score < 50 {
		fmt.Println("GRADE F")
	}
}
