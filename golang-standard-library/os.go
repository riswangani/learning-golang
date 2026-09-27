package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	hostame, err := os.Hostname()
	if err == nil {
		fmt.Println(hostame)
	} else {
		fmt.Println("Error:", err.Error())
	}
}
