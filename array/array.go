package main

import "fmt"

var productName [4]string

func main() {
	productName[0] = "Macbook"
	productName[1] = "iPad"
	productName[2] = "iPhone"
	productName[3] = "Airpods"

	price := [4]float32{40000, 3000, 20000, 2000}

	fmt.Println(productName)
	fmt.Println(price)
}
