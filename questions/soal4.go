package questions

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct {
	Address string
}

func (e *EmailNotifier) Send(msg string) {
	fmt.Printf("Sending Email to %s: %s\n", e.Address, msg)
}

type SMSNotifier struct {
	Number string
}

func (s *SMSNotifier) Send(msg string) {
	fmt.Printf("Sending SMS to %s: %s\n", s.Number, msg)
}

func SendAll(addr[]Notifier, msg string) {
	for _, n := range addr {
		n.Send(msg)
	}
}