package main

import (
	"sync"
)

// Отдельный мьютекс для каждой реализации
type FibonacciCalculator struct {
	deadlockCache struct {
		sync.Mutex
		values map[int]int
	}
	doubleCheckCache struct {
		sync.Mutex
		values map[int]int
	}
	syncMapCache sync.Map
}

func NewFibonacciCalculator() *FibonacciCalculator {
	return &FibonacciCalculator{
		deadlockCache: struct {
			sync.Mutex
			values map[int]int
		}{values: make(map[int]int)},
		doubleCheckCache: struct {
			sync.Mutex
			values map[int]int
		}{values: make(map[int]int)},
	}
}

// Версия с deadlock
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

// Исправленная версия
func (fc *FibonacciCalculator) FibonacciWithDoubleCheck(n int) int {
	fc.doubleCheckCache.Lock()
	if val, found := fc.doubleCheckCache.values[n]; found {
		fc.doubleCheckCache.Unlock()
		return val
	}
	fc.doubleCheckCache.Unlock()

	if n <= 1 {
		return n
	}

	val := fc.FibonacciWithDoubleCheck(n-1) + fc.FibonacciWithDoubleCheck(n-2)

	fc.doubleCheckCache.Lock()
	fc.doubleCheckCache.values[n] = val
	fc.doubleCheckCache.Unlock()

	return val
}

// Версия с sync.Map
func (fc *FibonacciCalculator) FibonacciWithSyncMap(n int) int {
	if val, ok := fc.syncMapCache.Load(n); ok {
		return val.(int)
	}

	if n <= 1 {
		return n
	}

	val := fc.FibonacciWithSyncMap(n-1) + fc.FibonacciWithSyncMap(n-2)
	fc.syncMapCache.Store(n, val)
	return val
}
