package main

import (
	"fmt"
)

func main() {
	//create a map with int as values and string as keys
	vertices := make(map[string]int)
	vertices["aayush"] = 1
	vertices["mittal"] = 2
	vertices["laptop"] = 3
	fmt.Println(vertices) //printing in alphabetical order
	delete(vertices, "laptop")
	fmt.Println(vertices)
	fmt.Println(vertices["aayush"])

}
