package main

import (
	"fmt"
)

func main() {
	type student struct {
		name  string
		age   int
		marks int
	}

	stud := student{name: "aayush", age: 20, marks: 90}
	fmt.Println(stud)
	fmt.Println(stud.name)

}
