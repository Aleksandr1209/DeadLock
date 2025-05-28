package main

import (
	"fmt"
	"time"
	"deadlock/taskFibonacci/problem"
	"deadlock/taskFibonacci/solution"
)

func main() {
	// Deadlock версия
	fmt.Println("Демонстрация deadlock:")
	fcProblem := problem.NewCalculator()
	done := make(chan bool)

	go func() {
		fmt.Println("Результат:", fcProblem.FibonacciWithDeadlock(20))
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Неожиданно: deadlock не произошёл")
	case <-time.After(2 * time.Second):
		fmt.Println("Deadlock подтверждён")
	}

	// Исправленные версии
	fcSolution := solution.NewCalculator()
	fmt.Println("DoubleCheck:", fcSolution.FibonacciWithDoubleCheck(20))
	fmt.Println("SyncMap:", fcSolution.FibonacciWithSyncMap(20))
}