package main

import "fmt"

func main() {

	fmt.Println(factorial(10))
	fmt.Println(FactorialWithRecursive(10))
}

func factorial(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func FactorialWithRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * FactorialWithRecursive(value - 1)
	}
}