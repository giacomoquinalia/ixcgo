package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	in := make(chan int)
	out := make(chan int)

	go worker(1, in, out)
	go worker(2, in, out)
	go worker(3, in, out)

	go generator(in)

	for o := range out {
		fmt.Println(o)
	}
}

func worker(id int, in chan int, out chan int) {
	for {
		value := <-in
		fmt.Printf("Worker %d está processando o valor %d\n", id, value)
		time.Sleep(1 * time.Second)
		out <- value
	}
}

// END OMIT

func generator(in chan int) {
	for i := 0; i < 10; i++ {
		in <- i
	}
	close(in)
}
