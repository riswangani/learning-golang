package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// func RunHelloWorld() {
// 	fmt.Print("Hello Wolrd")
// }

// func TestCreateGoroutine(t *testing.T) {
// 	go RunHelloWorld()

// 	fmt.Println("Ups")

// 	time.Sleep(1 * time.Second)
// }

// func DisplayNumber(number int) {
// 	fmt.Println("Display ", number)
// }


// func TestCreateChannel(t *testing.T) {
// 	channel := make(chan string)

// 	defer close(channel)

// 	go func() {

// 		time.Sleep(2 * time.Second)
// 		channel <- "Hello World"
// 		fmt.Println("Selesai kirim data ke channel")
		
// 	}()

// 	data := <-channel
// 	fmt.Println(data)
	
// 	time.Sleep(5 * time.Second)
// 	}


func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello Gani"
	// fmt.Println("Selesai kirim data ke channel")
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	
	time.Sleep(5 * time.Second)
	
	}