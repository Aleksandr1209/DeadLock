package tests

import (
	"math/rand"
	"sync"
	"testing"
	"time"
	"deadlock/taskBank/problem"
)

func TestConcurrentSafety(t *testing.T) {
	bank := problem.Bank{}
	for i := 0; i < 10; i++ {
		bank.Accounts = append(bank.Accounts, &problem.Account{ID: i, Balance: 1000})
	}

	var wg sync.WaitGroup
	success := make(chan bool, 1)

	go func() {
		<-time.After(5 * time.Second)
		success <- false
	}()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			from := rand.Intn(len(bank.Accounts))
			to := rand.Intn(len(bank.Accounts))
			if from != to {
				bank.TransferDeadlock(from, to, 1)
			}
		}()
	}

	go func() {
		wg.Wait()
		success <- true
	}()

	if !<-success {
		t.Fatal("Тест не завершился за 5 секунд (вероятный deadlock)")
	}

	total := 0
	for _, acc := range bank.Accounts {
		total += acc.Balance
	}
	if total != 10*1000 {
		t.Errorf("Общий баланс нарушен: got %d, want %d", total, 10*1000)
	}
}