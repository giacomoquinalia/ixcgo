package main

import (
	"fmt"
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
	go func() { output <- <-input1 }()
	go func() { output <- <-input2 }()
	return output
}

func generator(msg string) <-chan string {
	input := make(chan string)
	go func() {
		for i := 0; ; i++ {
			input <- fmt.Sprintf("%s %d", msg, i)
			// time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			time.Sleep(time.Second)
		}
	}()
	return input
}

// END OMIT
