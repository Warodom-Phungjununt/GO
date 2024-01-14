package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	CourseID   int     `json:"id"`
	CourseName string  `json:"name"`
	Price      float64 `json:"price"`
	Instructor string  `json:"instructor"`
}

func main() {
	e := employee{}
	err := json.Unmarshal([]byte(`{"ID":101,"EmployeeName":"Warodom","Tel":"0900000000","Email":"warodom.phu@gmail.com"}`), &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)
}
