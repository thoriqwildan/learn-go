package main

import "fmt"

func main() {
	bt := BankTransfer{AccountNumber: "123456789"}
	ew := EWallet{WalletID: "987-654-321"}

	ProsesPembayaran(bt, 100000)
	ProsesPembayaran(ew, 50000)
}

type PaymentMethod interface {
	Pay(amount float64)
}

type BankTransfer struct {
	AccountNumber string
}

func (b BankTransfer) Pay(amount float64) {
	fmt.Printf("Membayar Rp%.2f lewat Bank Transfer ke %s\n", amount, b.AccountNumber)
}

type EWallet struct {
	WalletID string
}

func (e EWallet) Pay(amount float64) {
	fmt.Printf("Membayar Rp%.2f lewat E-Wallet ID %s\n", amount, e.WalletID)
}

func ProsesPembayaran(method PaymentMethod, amount float64) {
	fmt.Println("Memproses pembayaran...")
	method.Pay(amount)
	fmt.Println("Pembayaran berhasil!")
}