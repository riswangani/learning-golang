package main

import (
	"errors"
	"fmt"
)


func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh 0")
	}
	return nilai / pembagi, nil
}

func main() {
	result, err := Pembagian(10, 2)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Result:", result)
	}
}