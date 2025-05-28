package problem

import (
	"sync"
	"time"
)

type Account struct {
	ID      int
	Balance int
	mu      sync.Mutex
}

type Bank struct {
	Accounts []*Account
}

func (b *Bank) TransferDeadlock(fromID, toID, amount int) {
	from := b.Accounts[fromID]
	to := b.Accounts[toID]

	from.mu.Lock()
	time.Sleep(time.Microsecond * 100)
	to.mu.Lock()

	from.Balance -= amount
	to.Balance += amount

	to.mu.Unlock()
	from.mu.Unlock()
}