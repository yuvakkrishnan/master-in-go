package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func init() {
	fmt.Println("Origin")

	balance = 1000
}
func deposit(amt int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Deposit %d to with balance: %d\n", amt, balance)
	balance += amt
	mutex.Unlock()
	wg.Done()
}

func withdraw(amt int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Withdrawing %d from account with balance %d\n", amt, balance)
	balance -= amt
	mutex.Unlock()
	wg.Done()
}

func main() {
	fmt.Println("Go Mutex Example")

	var wg sync.WaitGroup
	wg.Add(2)
	go withdraw(900, &wg)
	go deposit(1000, &wg)
	wg.Wait()
	fmt.Printf("New Balance%d\n", balance)
}
