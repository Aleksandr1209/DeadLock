package solution

import "sync"

type Account struct {
	ID      int
	Balance int
	mu      sync.Mutex
}

type Bank struct {
	Accounts []*Account
}

func (b *Bank) TransferCorrect(fromID, toID, amount int) {
	first, second := fromID, toID
	if fromID > toID {
		first, second = second, first
	}

	b.Accounts[first].mu.Lock()
	b.Accounts[second].mu.Lock()

	b.Accounts[fromID].Balance -= amount
	b.Accounts[toID].Balance += amount

	b.Accounts[second].mu.Unlock()
	b.Accounts[first].mu.Unlock()
}