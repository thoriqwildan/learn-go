package main

import "fmt"

func main() {
	var names = [...]string{
		"Thoriq",
		"Wildan",
		"Abdurrosyid",
		"Suka",
		"makan",
		"ayam",
		"dan",
		"juga",
		"nasi",
		"Goreng",
	}
	fmt.Println(names)
	iniSlice := names[2:6]
	fmt.Println(iniSlice)
}