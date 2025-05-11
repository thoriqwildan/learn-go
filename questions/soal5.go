package questions

import "fmt"

type PaymentMethod interface {
	ProcessPayment(amount float64) string
}

type CreditCard struct {
	CardNumber, HolderName string
}

func (c *CreditCard) ProcessPayment(amount float64) string {
	return "Pembayaran dengan kartu kredit " + c.CardNumber + " berhasil!"
}

type BankTransfer struct {
	BankAccount, BankName string
}

func (b *BankTransfer) ProcessPayment(amount float64) string {
	return "Pembayaran dengan kartu kredit " + b.BankAccount + " berhasil!"
}

func MakePayment(method PaymentMethod, amount float64) {
	fmt.Println(method.ProcessPayment(amount))
}