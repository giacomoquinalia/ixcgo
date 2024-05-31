package main

import "fmt"

// START OMIT
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go divideByTwo(ch1, ch2)
	go incrementByFive(ch2, ch3)

	ch1 <- 20
	fmt.Println(<-ch3)
}

func divideByTwo(in chan int, out chan int) {
	n := <-in
	n = n / 2
	out <- n
}

func incrementByFive(in chan int, out chan int) {
	n := <-in
	n = n + 5
	out <- n
}

// END OMIT
