package main

import "fmt"

func main() {
	// names := [3]string{"Thoriq", "Wildan", "Abdurrosyid"}

	// for index, name := range names {
	// 	fmt.Println("Index :", index, "Name :", name)
	// }
	
	// for _, k := range names {
	// 	fmt.Println(k)
	// }

	angka := map[int]string{ 1: "satu", 2: "dua", 3: "tiga", 4: "empat", 5: "lima" }

	for key, value := range angka {
		fmt.Printf("key %s: with Value %d\n", key, value)
	}
}