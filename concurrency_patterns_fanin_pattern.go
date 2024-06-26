package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func main() {
	output := fanIn(generator("Foo"), generator("Bar"))
	for {
		fmt.Println(<-output)
	}
}

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

// END OMIT

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
