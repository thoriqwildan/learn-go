package main

import "fmt"

func main() {
	Describe(3.14)
}

func Describe(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Println(len(v))
	case int:
		fmt.Println(v * 2)
	case bool:
		fmt.Println(v)
	default:
		fmt.Println("Tipe tidak dikenal")
	}
}