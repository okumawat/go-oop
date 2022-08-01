package model

type CreditCard struct {
	balance float64
}

func (c *CreditCard) GetBalance() float64 {
	return c.balance
}

func (c *CreditCard) ProcessPayment(payment float64) bool {
	c.balance += payment
	return true
}
