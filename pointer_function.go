package main

import "fmt"

func main() {
	bandung := Address{"Bandung", "Jawa Barat", "Singapura"}

	ChangeCountryToIndonesia(&bandung)
	fmt.Println(bandung)
}

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}