package main

import "fmt"

func main() {
	bl := func(name string) bool {
		return name == "anjing"
	}

	registerUser("anjing", bl)
}

type Blaclist func(string) bool

func registerUser(name string, blacklist Blaclist) {
	if blacklist(name) {
		fmt.Println("You're Blocked!")
	} else {
		fmt.Println("Welcome, " + name)
	}
}