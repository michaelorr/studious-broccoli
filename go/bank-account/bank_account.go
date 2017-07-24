package account

import "sync"

type Account struct {
	closed  bool
	balance int64
	mu      *sync.RWMutex
}

func Open(deposit int64) *Account {
	if deposit < 0 {
		return nil
	}
	return &Account{
		closed:  false,
		balance: deposit,
		mu:      &sync.RWMutex{},
	}
}

func (acc *Account) Close() (payout int64, ok bool) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	if acc.closed {
		return 0, false
	}
	acc.closed = true
	defer func() { acc.balance = 0 }()
	return acc.balance, true
}

func (acc *Account) Balance() (balance int64, ok bool) {
	acc.mu.RLock()
	defer acc.mu.RUnlock()
	if acc.closed {
		return 0, false
	}
	return acc.balance, true
}

func (acc *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	if acc.closed || acc.balance < -amount {
		return 0, false
	}
	acc.balance += amount
	return acc.balance, true
}
