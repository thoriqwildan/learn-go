package main

import "fmt"

func main() {
	name := "Joko"

	switch name {
	case "Thoriq":
		fmt.Println("Hallo Thoriq!")
	case "Joko":
		fmt.Println("Hallo Joko!")
		fallthrough
	case "Anwar":
		fmt.Println("Hallo Anwar!")
	default:
		fmt.Println("Hi, Boleh kenalan?")
	}
}