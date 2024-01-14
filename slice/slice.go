package main

import "fmt"

var courseName []string

func main() {
	var courseName []string
	courseName = []string{"Java", "Python"}
	fmt.Println(courseName)
	courseName = append(courseName, "C")
	fmt.Println(courseName)
	courseName = append(courseName, "C", "C#", "HTML", "CSS", "Javascript")
	fmt.Println(courseName)

	courseWeb := courseName[4:7]
	fmt.Println(courseWeb)
	courseWeb = courseName[:4]
	fmt.Println(courseWeb)
}
