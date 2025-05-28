package solution

import "sync"

type FibonacciCalculator struct {
	doubleCheckCache struct {
		sync.Mutex
		values map[int]int
	}
	syncMapCache sync.Map
}

func NewCalculator() *FibonacciCalculator {
	return &FibonacciCalculator{
		doubleCheckCache: struct {
			sync.Mutex
			values map[int]int
		}{values: make(map[int]int)},
	}
}

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