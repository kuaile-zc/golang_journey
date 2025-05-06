package sync_lock

import "sync"

var (
	lock    sync.Mutex
	balance float64
)

func Deposit(amount float64) {
	lock.Lock()
	defer lock.Unlock()
	balance += amount
}

func Balance() float64 {
	lock.Lock()
	defer lock.Unlock()
	return balance
}
