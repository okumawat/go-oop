package controller

type Payment interface {
	ProcessPayment(payment float64) bool
}
