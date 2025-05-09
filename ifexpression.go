package main

import (
	"fmt"
	"strconv"
)

func main() {
	if num, err := strconv.Atoi("143"); err == nil {
		fmt.Println("Angka : ", num)
	} else {
		fmt.Println("Gagal konversi : ", err)
	}
}