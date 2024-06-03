package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func main() {
	ch := fanIn(generator("Foo"), generator("Bar"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func fanIn(input1, input2 <-chan string) <-chan string {
	output := make(chan string)
	go func() {
		for {
			select {
			case v := <-input1:
				output <- v
			case v := <-input2:
				output <- v
			}
		}
	}()
	return output
}

// END OMIT

func generator(msg string) <-chan string {
	input := make(chan string)
	go func() {
		for i := 0; ; i++ {
			input <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return input
}
