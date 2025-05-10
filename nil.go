package main

import "fmt"

func main() {
	newMap := NewMap("s")

	if newMap == nil {
		fmt.Println("Map is nil")
	} else {
		fmt.Println("Map is not nil")
	}
}


func NewMap(name string) map[string] string {
	if name == "" {
		return nil
	} else {
		return map[string] string{
			"name": name,
		}
	}
}