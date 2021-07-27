package main

import (
	"fmt"
)

func main() {
	i := 6
	fmt.Println(&i)
	fmt.Println(i)
	inc(&i)
	fmt.Println(i)

}
func inc(x *int) {
	*x++
}
