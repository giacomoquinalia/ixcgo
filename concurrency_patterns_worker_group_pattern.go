package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// START OMIT
func main() {
	var sum atomic.Int64
	var wg sync.WaitGroup
	in := make(chan int64)
	out := make(chan int64)

	wg.Add(3)
	go worker(1, in, out, &wg)
	go worker(2, in, out, &wg)
	go worker(3, in, out, &wg)

	go func() {
		wg.Wait()
		close(out)
	}()

	go generator(in)
	for o := range out {
		sum.Add(o)
	}
	fmt.Println(sum.Load())
}

// END OMIT

func worker(id int, in <-chan int64, out chan<- int64, wg *sync.WaitGroup) {
	for v := range in {
		fmt.Printf("Worker %d está processando o valor %d\n", id, v)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		out <- v
	}
	wg.Done()
}

func generator(in chan int64) {
	defer close(in)
	for i := 0; i < 10; i++ {
		in <- int64(i)
	}
}
