package main

import "fmt"

func concat(s, t string) string{
	return s + t
}

func function(){
	fmt.Println(concat("Go ", "is Amazing"))
}