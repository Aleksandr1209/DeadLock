package main

import (
	"fmt"
	"time"
	"deadlock/taskFibonacci/problem"
)

func main() {
	// Deadlock версия
	fmt.Println("=== ДЕМОНСТРАЦИЯ РАБОТЫ ЗАДАЧИ С ЧИСЛАМИ ФИБОНАЧЧИ ===")
	fcProblem := problem.NewCalculator()
	done := make(chan bool)

	go func() {
		fmt.Println("Результат:", fcProblem.FibonacciWithDeadlock(20))
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Программа успешно завершила работу")
	case <-time.After(3 * time.Second):
		fmt.Println("Обнаружен Deadlock")
	}
}