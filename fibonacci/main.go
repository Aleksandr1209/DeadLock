package main

import (
	"fmt"
	"time"
)

func main() {
	fc := NewFibonacciCalculator()

	fmt.Println("1. Демонстрация deadlock (нажмите Ctrl+C чтобы пропустить):")
	deadlockDetected := make(chan bool)

	go func() {
		fmt.Println("Результат:", fc.FibonacciWithDeadlock(20))
		deadlockDetected <- false
	}()

	select {
	case <-deadlockDetected:
		fmt.Println("Неожиданно: deadlock не произошёл")
	case <-time.After(2 * time.Second):
		fmt.Println("2. Deadlock подтверждён (функция не завершилась)")
	}

	fmt.Println("\n3. Исправленная версия с двойной проверкой:")
	fmt.Println("Результат:", fc.FibonacciWithDoubleCheck(20))

	fmt.Println("\n4. Версия с sync.Map:")
	fmt.Println("Результат:", fc.FibonacciWithSyncMap(20))
}
