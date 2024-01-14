package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	Id          int
	EmplyeeName string
	Tel         string
	Email       string
}

func main() {
	data, _ := json.Marshal(&employee{101, "Warodom Phungjununt", "0900000000", "warodom.phu@gmail.com"})
	fmt.Println(string(data))
}
