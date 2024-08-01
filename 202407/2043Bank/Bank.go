package main

func main() {

}

type Bank struct {
	Balance []int64
	N       int
}

func Constructor(balance []int64) Bank {
	return Bank{Balance: balance, N: len(balance)}

}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1 <= 0 || account1 > this.N {
		return false
	}
	if account1 <= 0 || account2 > this.N {
		return false
	}

	if this.Balance[account1-1] < money {
		return false
	}
	this.Balance[account1-1] -= money
	this.Balance[account2-1] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if account <= 0 || account > this.N {
		return false
	}
	this.Balance[account-1] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if account <= 0 || account > this.N {
		return false
	}
	if this.Balance[account-1] < money {
		return false
	}
	this.Balance[account-1] -= money
	return true
}
