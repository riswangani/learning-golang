package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	// ini adalah pass by value
	// address1 := Address{
	// 	City: "Jakarta",
	// 	Province: "DKI Jakarta",
	// 	Country: "Indonesia",
	// }

	// address2 := address1
	// address2.City = "Bandung"
	

	// fmt.Println(address1)
	// fmt.Println(address2)

  // Pointer
	// ini adalah pass by reference 
	var address1 Address = Address{
		City: "Jakarta",
		Province: "DKI Jakarta",
		Country: "Indonesia",
	}

	var address2 *Address = &address1 // pointer
	//address2.City = "Bandung"
	

	fmt.Println(address1)
	fmt.Println(address2)

	// mengubah address2 maka address1 juga akan berubah
	// karena address2 adalah pointer dari address1
	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)

	

	
}