package main

import "fmt"

type Person struct {
	name string
	age int
	address string
}

func (p Person) sayHello() {
	fmt.Println("Hello", p.name)
}

func main() {
	var orang Person
	orang.name = "Cahyo"
	orang.age = 187
	orang.address = "Jatimulyo"

	thoriq := Person{
		name: "Thoriq",
		age: 18,
		address: "Ngaglik",
	}

	fmt.Println(orang)
	fmt.Println(thoriq)

	thoriq.sayHello()
}