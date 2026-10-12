package main

import (
	"fmt"

	"github.com/riswangani/app-say-hello/users"
	go_say_hello "github.com/riswangani/go-say-hello/v2"
)

func main() {
	// Memanggil fungsi dari package users (lokal)
	fmt.Println(users.SayHello("Riswan"))

	// Memanggil modul github.com/riswangani/go-say-hello/v2
	fmt.Println(go_say_hello.SayHello("GANI BAIK"))
}
