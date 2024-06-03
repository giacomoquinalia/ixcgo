package main

import (
	"fmt"
	"time"
)

func worker(tickets <-chan Ticket, work <-chan Work) {
	for w := range work {
		<-tickets // espera por um ticket
		w()       // executa um trabalho
	}
}

func boxOffice(tickets chan<- Ticket, timeout time.Duration, nTickets int) {
	for {
		for i := 0; i < nTickets; i++ {
			tickets <- Ticket(i)
		}

		// espera até que mais tickets possam ser emitidos
		<-time.After(timeout)
	}
}

// START OMIT
type Work func()
type Ticket int

func main() {
	tickets := make(chan Ticket)
	works := make(chan Work)

	go boxOffice(tickets, 1*time.Second, 10)
	go worker(tickets, works)

	for i := 0; i <= 30; i++ {
		works <- func() {
			fmt.Println("processando ticket")
		}
		fmt.Println("Work ", i, " sent")
	}

	close(works)
	close(tickets)
}

// END OMIT
