package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountry(address *Address) {
	
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{}
		

	changeCountry(&address1)

	fmt.Println(address1)
}