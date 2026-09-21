package main

import "fmt"


type Man struct {
	Name string
}

func(man *Man) Mister(){
	man.Name = "Mr. " + man.Name
}

func main() {
	gani := Man{"Gani"}
	gani.Mister()
	fmt.Println(gani.Name)
}