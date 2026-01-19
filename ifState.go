package main

import "fmt"

func ifState() {
	email := "aditya@123"

	if length := len(email); length < 1 {
		fmt.Println("Email Invalid")
	} else {
		fmt.Println("Valid Email")
	}
}
