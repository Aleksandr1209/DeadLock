package tests

import (
	"testing"
	"time"
	"deadlock/taskFibonacci/problem"
)

func TestCorrectnessAfterFix(t *testing.T) {
	fc := problem.NewCalculator()

	if res := fc.FibonacciWithDeadlock(0); res != 0 {
		t.Errorf("Fib(0): ожидалось 0, получено %d", res)
	}
	if res := fc.FibonacciWithDeadlock(1); res != 1 {
		t.Errorf("Fib(1): ожидалось 1, получено %d", res)
	}

	testCases := []struct {
		n, expected int
	}{
		{5, 5},
		{10, 55},
		{15, 610},
	}

	for _, tc := range testCases {
		done := make(chan bool)
		var result int

		go func(n int) {
			result = fc.FibonacciWithDeadlock(n)
			done <- true
		}(tc.n)

		select {
		case <-done:
			if result != tc.expected {
				t.Errorf("Fib(%d): ожидалось %d, получено %d", tc.n, tc.expected, result)
			}
		case <-time.After(1 * time.Second):
			t.Errorf("Fib(%d): функция не завершилась (deadlock)", tc.n)
		}
	}
}