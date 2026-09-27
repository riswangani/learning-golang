package helper

import (
	"fmt"
	"testing"
	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// func TestHelloWorld(t *testing.T) {
// 	result := HelloWorld("Budi")
// 	if result != "Hello Budi" {
// 		t.Fail()
// 		// panic("Result is not 'Hello Budi'")
// 	}

// 	fmt.Println("Test Hello World Selesai")
// }


// func TestHelloWorldGani(t *testing.T) {
// 	result := HelloWorld("Gani")
// 	if result != "Hello Gani" {
// 		//t.Errorf("Expected 'Hello Gani' but got '%s'", result)
// 		// panic("Result is not 'Hello Ganif'")
// 		t.FailNow()
// 	}

// 	fmt.Println("Test Hello World Gani Selesai")
// }


// func TestHelloWorld(t *testing.T) {
// 	result := HelloWorld("Budi")
// 	if result != "Hello Budi" {
// 		t.Error("Result must be 'Hello Budi'")
// 		// panic("Result is not 'Hello Budi'")
// 	}

// 	fmt.Println("Test Hello World Selesai")
// }


// func TestHelloWorldGani(t *testing.T) {
// 	result := HelloWorld("Gani")
// 	if result != "Hello Gani" {
// 		//t.Errorf("Expected 'Hello Gani' but got '%s'", result)
// 		// panic("Result is not 'Hello Ganif'")
// 		t.Fatal("Result must be 'Hello Gani'")

// 	}

// 	fmt.Println("Test Hello World Gani Selesai")


// }

// pake testify assertion
// func TestHelloWorldAssert(t *testing.T) {
// 	result := HelloWorld("Budi")
// 	assert.Equal(t, "Hello Budi", result, "Result must be 'Hello Budi'")

// 	fmt.Println("Test Hello World Assert Selesai")	
// }

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Budi")
	require.Equal(t, "Hello Budi", result, "Result must be 'Hello Budi'")

	fmt.Println("Test Hello World Require Selesai")	
}