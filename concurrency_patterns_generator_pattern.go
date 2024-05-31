package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	ch := make(chan int)

	go generator(ch)

	for n := range ch {
		fmt.Println(n)
	}
}

func generator(ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		ch <- i
	}
	close(ch)
}

// END OMIT
