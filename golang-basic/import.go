package main

import (
	"fmt"
	"golang-basic/helper"
	"golang-basic/database"
	_"golang-basic/internal"
	
)

func main() {
	resutlt := helper.SayHello("Gani")
	fmt.Println(resutlt)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error karena pake hurup kecil

	fmt.Println(database.GetConnection())
}