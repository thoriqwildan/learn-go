package main

import "fmt"

func main() {
	// nama := []string{"Thoriq", "Wildan", "Cahyo", "Eko", "Budi", "Yasin"}

	// haloNama("Halo", nama...)

	// hasil := hitung(10, 5, tambah)
	// fmt.Println(hasil)

	sayHelloWithFilter("anjing", filter)
}

func haloNama(msg string, names ...string) {
	for _, name := range names {
		fmt.Println(msg, name)
	}
}

func hitung(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func tambah(a, b int) int {
	return a + b
}

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func filter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}