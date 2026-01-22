package main

import "fmt"

// Interface
type Shape interface{
	Area() int           // the method in the parameter should also have a return datatype
}

type Perimeter interface{
	Perimeter() int
}

type Geometry interface{
	Shape
	Perimeter
}


type recta struct{
	width int
	length int
}

func (r recta) Area() int{
	return r.width * r.length
}

func(r recta) Perimeter() int{
	return 2 *(r.width + r.length)
}

// type circle struct{
// 	radius int
// }

// func (c circle) Area() int{
// 	return 3*c.radius*c.radius
// }

func describeShape(g Geometry) {
	fmt.Println("Area of the appropriate shape is:", g.Area())
	fmt.Println("Perimeter of the appropriate shape is:", g.Perimeter())
}

func interfac(){
	r := recta{width: 5, length: 4}
	// c:= circle{5}

	describeShape(r)
}