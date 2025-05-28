package tests

import (
	"testing"
	"time"
	"deadlock/taskFibonacci/problem"
	"deadlock/taskFibonacci/solution"
)

func TestFibonacciWithDeadlock(t *testing.T) {
	fc := problem.NewCalculator()

	done := make(chan bool)
	go func() {
		fc.FibonacciWithDeadlock(10)
		done <- true
	}()

	select {
	case <-done:
		t.Error("Ожидался deadlock, но функция завершилась")
	case <-time.After(1 * time.Second):
	}
}

func TestFibonacciWithDoubleCheck(t *testing.T) {
	fc := solution.NewCalculator()

	results := make(chan int, 1)
	go func() {
		results <- fc.FibonacciWithDoubleCheck(20)
	}()

	select {
	case res := <-results:
		t.Logf("Fib(20) = %d", res)
	case <-time.After(2 * time.Second):
		t.Error("Функция не завершилась в течение 2 секунд")
	}
}

func TestFibonacciWithSyncMap(t *testing.T) {
	fc := solution.NewCalculator()

	results := make(chan int, 1)
	go func() {
		results <- fc.FibonacciWithSyncMap(20)
	}()

	select {
	case res := <-results:
		t.Logf("Fib(20) = %d", res)
	case <-time.After(2 * time.Second):
		t.Error("Функция не завершилась в течение 2 секунд")
	}
}