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

	go boxOffice(tickets, 1*time.Second, 2)
	go worker(tickets, works)

	for i := 0; i <= 6; i++ {
		fmt.Printf("Enviando trabalho %d\n", i)
		works <- func() {
			fmt.Printf("Processando trabalho %d\n", i)
		}
	}

	close(works)
	close(tickets)
}

// END OMIT
