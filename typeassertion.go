package main

import "fmt"

func main() {
	// result := random().(string)
	// fmt.Println(result)

	// resultInt := random().(int)
	// fmt.Println(resultInt)

	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("String:", value)
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float64:", value)
	case bool:
		fmt.Println("Boolean:", value)
	default:
		fmt.Println("Unknown type")
	}
}

func random() interface{} {
	return "OK"
}

