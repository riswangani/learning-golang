package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("Gani")

	if data == nil {
		fmt.Println("Map is nil")
	} else {
		fmt.Println(data["name"])
	}
}