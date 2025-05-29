package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"deadlock/taskBank/problem"
	"deadlock/taskBank/solution"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	const numAccounts = 100
	const numTransactions = 10000

	fmt.Println("=== ДЕМОНСТРАЦИЯ РАБОТЫ С БАНКОВСКИМИ ПЕРЕВОДАМИ ===")

	bankDeadlock := problem.Bank{}
	for i := 0; i < numAccounts; i++ {
		bankDeadlock.Accounts = append(bankDeadlock.Accounts, &problem.Account{ID: i, Balance: 1000})
	}
	runTransactions(&bankDeadlock, numTransactions, true)
}

func runTransactions(bank interface{}, count int, useDeadlock bool) {
	var wg sync.WaitGroup
	success := make(chan bool, 1)

	go func() {
		<-time.After(5 * time.Second)
		success <- false
	}()

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			from := rand.Intn(100)
			to := rand.Intn(100)
			if from == to {
				return
			}

			switch b := bank.(type) {
			case *problem.Bank:
				b.TransferDeadlock(from, to, 1)
			case *solution.Bank:
				b.TransferCorrect(from, to, 1)
			}
		}()
	}

	go func() {
		wg.Wait()
		success <- true
	}()

	if <-success {
		fmt.Println("Все переводы завершены успешно!")
	} else {
		fmt.Println("Обнаружен deadlock!")
	}
}