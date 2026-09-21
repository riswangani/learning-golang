package main

import "fmt"

func random() any {
	return true
}

func main() {
	result := random()

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}