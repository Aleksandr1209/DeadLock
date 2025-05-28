package problem

import "sync"

type FibonacciCalculator struct {
	deadlockCache struct {
		sync.Mutex
		values map[int]int
	}
}

func NewCalculator() *FibonacciCalculator {
	return &FibonacciCalculator{
		deadlockCache: struct {
			sync.Mutex
			values map[int]int
		}{values: make(map[int]int)},
	}
}

func (fc *FibonacciCalculator) FibonacciWithDeadlock(n int) int {
	fc.deadlockCache.Lock()
	defer fc.deadlockCache.Unlock()

	if val, found := fc.deadlockCache.values[n]; found {
		return val
	}

	if n <= 1 {
		return n
	}

	val := fc.FibonacciWithDeadlock(n-1) + fc.FibonacciWithDeadlock(n-2)
	fc.deadlockCache.values[n] = val
	return val
}