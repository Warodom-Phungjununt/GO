package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {
	employeeList := [5]employee{}
	employeeList[0] = employee{
		employeeID:   "101",
		employeeName: "Warodom",
		phone:        "01000001",
	}
	employeeList[1] = employee{
		employeeID:   "102",
		employeeName: "Warodam",
		phone:        "01000002",
	}
	employeeList[2] = employee{
		employeeID:   "103",
		employeeName: "Warodem",
		phone:        "01000003",
	}
	employeeList[3] = employee{
		employeeID:   "104",
		employeeName: "Warodim",
		phone:        "01000004",
	}
	employeeList[4] = employee{
		employeeID:   "105",
		employeeName: "Warodum",
		phone:        "01000005",
	}
	fmt.Println("employee = ", employeeList)
}
