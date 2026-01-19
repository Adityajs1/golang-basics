package main

import "fmt"

func typecast(){
	ageInt := 2.6
    ageYears := int(ageInt)

	fmt.Println("The age of the account is ", ageYears)
}