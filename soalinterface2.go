package main

import "fmt"

func main() {
	notifiers := []Notifier{
		EmailNotifier{EmailAddress: "wildan@gmail.com"},
		SMSNotifier{PhoneNumber: "08213434254"},
	}

	for _, n := range notifiers {
		n.Notify("Hello, this is a notification!")
	}
}

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct {
	EmailAddress string
}

func (e EmailNotifier) Notify(message string) {
	fmt.Println("Email sent to", e.EmailAddress, "with message:", message)
}

type SMSNotifier struct {
	PhoneNumber string
}

func (s SMSNotifier) Notify(message string) {
	fmt.Println("SMS sent to", s.PhoneNumber, "with message:", message)
}