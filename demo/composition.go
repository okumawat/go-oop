package demo

import "fmt"

type Account struct {
	balance int
}

func (a *Account) ShowBalance() {
	fmt.Println("Balance:", a.balance)
}

//Composition
type SavingAccount struct {
	Account
}

type CreditAccount struct {
	Account
}

type HybridAccount struct {
	SavingAccount
	CreditAccount
}

func (h *HybridAccount) ShowBalance() {
	h.SavingAccount.ShowBalance()
}

func CompositionDemo() {
	ha := HybridAccount{}
	ha.ShowBalance()
}
