package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2) //used so that channels doesnt block program

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)

}
