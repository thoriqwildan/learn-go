package main

import (
	"fmt"
	"latsol/questions"
)

func main() {
	// Q1()
	// Q2()
	// Q3()
	// Q4()
	Q5()
}

func Q1() {
	a := 5
	b := 10
	questions.Swap(&a, &b)
	fmt.Println(a, b) // Output: 10 5
}

func Q2() {
	data := []int{70, 80, 90, 100}
	bonus := 5
	questions.AddBonus(data, &bonus)
	fmt.Println(data) // Output: [75 85 95 105]
}

func Q3() {
	u := questions.User{Name: "Anon", Age: 20}
	questions.UpdateProfile(&u, "John", 30)
	fmt.Println(u.Name, u.Age) // Output: John 30
}

func Q4() {
	notifiers := []questions.Notifier {
		&questions.EmailNotifier{"user@example.com"},
		&questions.SMSNotifier{"37582374522"},
	}

	questions.SendAll(notifiers, "Hello World")
}

func Q5() {
	payment1 := &questions.CreditCard{"1234-5678-9876-5432", "John Doe"}
	payment2 := &questions.BankTransfer{"123-456-789", "Bank A"}

	questions.MakePayment(payment1, 500.0) // Output: Pembayaran dengan kartu kredit 1234-5678-9876-5432 berhasil!
	questions.MakePayment(payment2, 1000.0) // Output: Pembayaran melalui transfer bank 123-456-789 berhasil!

}