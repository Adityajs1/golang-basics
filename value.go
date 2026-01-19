package main

import "fmt"

func value(){
	x := 5
	increment(x)
	fmt.Println(x)
}

func increment(x int){
	x++
}