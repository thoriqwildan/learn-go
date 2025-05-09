package main

import "fmt"

func main() {
	runApplication(true)
}

func logging() {
	fmt.Println("Semua task selesai")
	msg := recover()
	fmt.Println("Error di :", msg)
}

func runApplication(error bool) {
	defer logging() // Penutup function, tetap berjalan walaupun function mengalami crash
	if error {
		panic("ERROR!!!") // panic digunakan untuk memberitahu bahwa function error
	}
	fmt.Println("Running App")
}