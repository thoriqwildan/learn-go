package main

import "fmt"

func main() {
	price := 50000
	fmt.Println("Harga awal: Rp", price)

	fmt.Println("Tanpa diskon: Rp", PrintFinalPrice(float64(price), NoDiscount{}))
	fmt.Println("Diskon 10%: Rp", PrintFinalPrice(float64(price), PercentageDiscount{}))
	fmt.Println("Diskon Rp15000: Rp", PrintFinalPrice(float64(price), FlatDiscount{}))
}

type Discount interface {
	Calculate(price float64) float64
}

type NoDiscount struct {}
func (d NoDiscount) Calculate(price float64) float64 {
	return price
}

type PercentageDiscount struct {}
func (d PercentageDiscount) Calculate(price float64) float64 {
	return price * 0.9
}

type FlatDiscount struct {}
func (d FlatDiscount) Calculate(price float64) float64 {
	if price < 15000 {
		return price
	} else {
		return price - 15000
	}
}

func PrintFinalPrice(price float64, discount Discount) float64 {
	return discount.Calculate(price)
}