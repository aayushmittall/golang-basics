package main

import (
	"fmt"
	"time"
)

func main() {
	go count("aayush")
	count("mittal")
}

func count(thing string) {
	for i := 1; i < 10; i++ {
		fmt.Println(thing)
		time.Sleep(time.Millisecond * 500)
	}
}
