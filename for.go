package main

import (
	"fmt"
)

func main() {
	arr := [7]int{1, 2, 3, 4, 5}
	m := make(map[string]string)
	m["aayush"] = "mittal"
	m["mittal"] = "aayush"

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	for i, value := range arr {
		fmt.Println(i, value)
	}

	for key, value := range m {
		fmt.Println(key, value)
	}

}
