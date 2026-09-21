package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{
		City: "Jakarta",
		Province: "DKI Jakarta",
		Country: "Indonesia",
	}

	address2 := &address1 // pointer
	address2.City = "Bandung"

	// address2 = &Address{
	// 	City: "Jakarta",
	// 	Province: "DKI Jakarta",
	// 	Country: "Indonesia",
	// }

	*address2 = Address{
		City: "Jakarta",
		Province: "DKI Jakarta",
		Country: "Indonesia",
	}
	

	fmt.Println(address1)
	fmt.Println(address2)
}