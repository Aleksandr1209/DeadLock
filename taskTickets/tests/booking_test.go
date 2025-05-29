package tests

import (
	"testing"
	"time"

	"deadlock/taskTickets/problem"
)

func TestDeadlockScenario(t *testing.T) {
	bs := problem.NewBookingSystem(2, 10)

	done := make(chan struct{})

	go func() {
		bs.Book(0, 1)
		done <- struct{}{}
	}()

	go func() {
		bs.Book(1, 0) 
		done <- struct{}{}
	}()

	select {
	case <-done:
		t.Log("Deadlock не обнаружен")
	case <-time.After(3 * time.Second):
		t.Error("Обнаружен deadlock")
	}
}

