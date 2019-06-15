package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("sheep", c)

	// for i := 1; i < 12; i++ {
	// 	msg := <-c
	// 	fmt.Println(msg)
	// }

	for {
		msg, open := <-c
		if !open {
			break
		}
		fmt.Println(msg)
	}

	// for msg := range c {
	// 	fmt.Println(msg)
	// }
}

func count(thing string, c chan string) {
	for i := 1; i < 10; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c) //without this there will be deadlock as in main thread for loop is infinie and it'll be wanting msg!!
}
