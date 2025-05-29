package main

import (
	"fmt"
	"time"
	"deadlock/taskTickets/problem"
)

func main() {
	bs := problem.NewBookingSystem(2, 10)

	fmt.Println("Simulating guaranteed deadlock...")

	done := make(chan struct{})

	go func() {
		fmt.Println("Goroutine 1: booking from 0 to 1")
		bs.Book(0, 1)
		fmt.Println("Goroutine 1: done")
		done <- struct{}{}
	}()

	go func() {
		fmt.Println("Goroutine 2: booking from 1 to 0")
		bs.Book(1, 0)
		fmt.Println("Goroutine 2: done")
		done <- struct{}{}
	}()

	select {
	case <-done:
		fmt.Println("Программа успешно завершилась")
	case <-time.After(5 * time.Second):
		fmt.Println("Обнаружен DeadLock")
	}
}

