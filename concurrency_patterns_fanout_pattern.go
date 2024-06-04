package main

import (
	"fmt"
	"sync"
	"time"
)

func publish(output chan<- int, value int, wg *sync.WaitGroup) {
	select {
	case output <- value:
	case <-time.After(1 * time.Second):
	}
	wg.Done()
}

func sequence(start, end int) <-chan int {
	input := make(chan int)
	go func() {
		for i := start; i <= end; i++ {
			input <- i
		}
		close(input)
	}()
	return input
}

func worker(id int, in <-chan int, wg *sync.WaitGroup) {
	for v := range in {
		fmt.Printf("Worker %d está processando o valor %d\n", id, v)
	}
	wg.Done()
}

// START OMIT
func main() {
	ch1, ch2 := make(chan int), make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(1, ch1, &wg)
	go worker(2, ch2, &wg)
	fanOut(sequence(1, 10), ch1, ch2)
	wg.Wait()
}
func fanOut(input <-chan int, outputs ...chan<- int) {
	var wg sync.WaitGroup
	for v := range input {
		wg.Add(len(outputs))
		for _, output := range outputs {
			go publish(output, v, &wg)
		}
		wg.Wait()
	}
	for _, output := range outputs {
		close(output)
	}
}

// END OMIT
