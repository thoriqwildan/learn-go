package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Pembagian(100, 0)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

func Pembagian(value int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		return value / pembagi, nil
	}
}