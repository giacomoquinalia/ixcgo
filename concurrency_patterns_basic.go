package main

import (
	"fmt"
)

// START OMIT
func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello, world!"
	}()

	fmt.Println(<-ch)
}

// END OMIT
