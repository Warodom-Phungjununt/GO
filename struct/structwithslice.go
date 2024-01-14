package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {
	employeeList := []employee{}
	employee1 := employee{
		employeeID:   "101",
		employeeName: "Warodom",
		phone:        "0800001",
	}
	employee2 := employee{
		employeeID:   "102",
		employeeName: "Warodam",
		phone:        "0800002",
	}
	employee3 := employee{
		employeeID:   "103",
		employeeName: "Warodem",
		phone:        "0800003",
	}

	employeeList = append(employeeList, employee1)
	employeeList = append(employeeList, employee2)
	employeeList = append(employeeList, employee3)

	fmt.Println("employee = ", employeeList)
}
