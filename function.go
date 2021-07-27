package main

import (
	"fmt"
)

func main() {
	result := sum(4, 3)
	fmt.Println(result)
}

func sum(x int, y int) int {
	return x + y
}
