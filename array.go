package main

import (
	"fmt"
)

func main() {
	var array [5]int //array fixed size
	array[1] = 2
	fmt.Println(array)

	arr := [5]int{1, 2, 3, 4, 5} // can't be appended
	fmt.Println(arr)

}
