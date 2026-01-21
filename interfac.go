package main

import "fmt"

// Interface
type Shape interface{
	Area() int           // the method in the parameter should also have a return datatype
}

type recta struct{
	width int
	length int
}

func (r recta) Area() int{
	return r.width * r.length
}

type circle struct{
	radius int
}

func (c circle) Area() int{
	return 3*c.radius*c.radius
}

func PrintArea(s Shape) {
	fmt.Println("Area of the appropriate shape is:", s.Area())
}

func interfac(){
	r := recta{5,4}
	c:= circle{5}

	PrintArea(r)
	PrintArea(c)
}