package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("aayush")
		wg.Done() //decrements the value of wg by 1
	}()
	wg.Wait() //program runs util it is zzero.
}

func count(thing string) {
	for i := 1; i < 10; i++ {
		fmt.Println(thing)
		time.Sleep(time.Millisecond * 500)
	}
}
