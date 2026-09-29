package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)


	func OnlyIn(channel chan<- string) {
		time.Sleep(2 * time.Second)
		channel <- "Hello Gani"
	}

	func OnlyOut(channel <-chan string) {
		data := <-channel
		fmt.Println(data)
	}


	func TestInOutChannel(t *testing.T) {
			channel := make(chan string)
			defer close(channel)

			go OnlyIn(channel)
			go OnlyOut(channel)
		
			time.Sleep(5 * time.Second)		
	}