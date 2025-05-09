package main

import "fmt"

func main() {
	SayHelloWithFilter("Wildan", anjing)
}

type Filter func(string) string

func SayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func anjing(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}