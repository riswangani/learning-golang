package golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	// Ambil value dari contextF
	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))

	fmt.Println(contextA.Value("b"))
}

// tanpa context cancel, goroutine akan terus berjalan walaupun sudah keluar dari loop
// func CreateCounter() chan int {

// 	destination := make(chan int)

// 	go func() {
// 		defer close(destination)
// 		couterter := 1
// 		for {
// 			destination <- couterter
// 			couterter++
// 		}
// 	}()

// 	return destination
// }

// func TestContextWithCancel(t *testing.T) {
// fmt.Println("Total Goroutine", runtime.NumGoroutine())

// destionation := CreateCounter()
// for n := range destionation {
// 	fmt.Println("Counter", n)
// 	if n == 10 {
// 		break
// 	}
// }

// fmt.Println("Total Goroutine", runtime.NumGoroutine())
// }
//

// dengan context cancel, goroutine akan berhenti ketika keluar dari loop

// func CreateCounter(ctx context.Context) chan int {

// 	destination := make(chan int)

// 	go func() {
// 		defer close(destination)
// 		couterter := 1
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				return
// 			default:
// 				destination <- couterter
// 				couterter++
// 			}
// 		}
// 	}()

// 	return destination
// }

// func TestContextWithCancel(t *testing.T) {
// fmt.Println("Total Goroutine", runtime.NumGoroutine())

// parent := context.Background()
// ctx, cancel := context.WithCancel(parent)

// destionation := CreateCounter(ctx)

// fmt.Println("Total Goroutine", runtime.NumGoroutine())

// for n := range destionation {
// 	fmt.Println("Counter", n)
// 	if n == 10 {
// 		break
// 	}
// }

// cancel() //mengirim sinyal cancel ke goroutine untuk berhenti

// time.Sleep(2* time.Second)

// fmt.Println("Total Goroutine", runtime.NumGoroutine())
// }

// // dengan context cancel dengan timeout, goroutine akan berhenti ketika keluar dari loop
// func CreateCounter(ctx context.Context) chan int {
// 	destination := make(chan int)

// 	go func() {
// 		defer close(destination)
// 		couterter := 1
// 		for {
// 			select {
// 			case <- ctx.Done():
// 				return
// 			default:
// 				destination <- couterter
// 				couterter++
// 				time.Sleep(1 * time.Second) // simulasi proses yang memakan waktu
// 			}
// 		}
// 	}()

// 	return destination
// }

// func TestContextWithTimeOut(t *testing.T) {
// 	fmt.Println("Total Goroutine", runtime.NumGoroutine())

// 	parent := context.Background()
// 	ctx, cancel := context.WithTimeout(parent, 5 * time.Second)
// 	defer cancel()

// 	destionation := CreateCounter(ctx)

// 	fmt.Println("Total Goroutine", runtime.NumGoroutine())

// 	for n := range destionation {
// 		fmt.Println("Counter", n)
// 	}


// 	time.Sleep(2 * time.Second)

// 	fmt.Println("Total Goroutine", runtime.NumGoroutine())
// }


// dengan context cancel dengan deadline, goroutine akan berhenti ketika keluar dari loop
func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		couterter := 1
		for {
			select {
			case <- ctx.Done():
				return
			default:
				destination <- couterter
				couterter++
				time.Sleep(1 * time.Second) // simulasi proses yang memakan waktu
			}
		}
	}()

	return destination
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5 * time.Second))
	defer cancel()

	destionation := CreateCounter(ctx)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	for n := range destionation {
		fmt.Println("Counter", n)
	}


	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
