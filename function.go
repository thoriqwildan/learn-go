package main

import "fmt"

func main() {
	first, _ := getFullName()
	hello := getHello(first)
	fmt.Println(hello)
}

// func sayHello() {
// 	fmt.Println("Hello World!")
// }

// func sayHello(name string) {
// 	fmt.Println("Hello,", name)
// }

func getHello(name string) string {
	return "Hello " + name
} 

func getFullName() (string, string) {
	return "Thoriq", "Wildan"
}