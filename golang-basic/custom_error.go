package main

import "fmt"

type validationError struct {
	Message string
}

func (e *validationError) Error() string {
	return e.Message
}

type notFoundError struct {
	Message string
}

func (e *notFoundError) Error() string {
	return e.Message
}


func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}

	if id != "gani" {
		return &notFoundError{"not found error"}
		}

	//oke

	return nil
}

func main() {
	err := SaveData("gani", nil)
	if err != nil {
		//terjadi error
		if validationError, ok := err.(*validationError); ok {
			fmt.Println("validation error:", validationError.Error())
		} else if notFoundError, ok := err.(*notFoundError); ok {
			fmt.Println("not found error:", notFoundError.Error())
		} else {
			fmt.Println("Unknown Error:", err.Error())
		}
	} else {
		fmt.Println("Success")
	}
}
