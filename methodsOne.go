package main

import "fmt"

// struct definition
type rect struct {
	width  int
	length int
}

// method on rect
func (r rect) area() int {
	return r.width * r.length
}

func methodsOne() {
	// creating rect object
	r := rect{
		width:  5,
		length: 10,
	}

	// calling method
	fmt.Println(r.area())
}
