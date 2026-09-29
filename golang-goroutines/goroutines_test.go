package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)


func RunHelloWorld() {
	fmt.Print("Hello Wolrd")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()

	fmt.Println("Ups")
	
	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display ", number)
}


func TestManyGoroutines(t *testing.T) {
	for i := range 100000 {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}