package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {
	employee1 := employee{
		employeeID:   "101",
		employeeName: "The.Dome",
		phone:        "09999999990",
	}
	fmt.Println("employee1 = ", employee1)
}
