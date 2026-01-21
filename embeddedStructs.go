package main

import "fmt"

// Base struct (will be embedded)
type Engine struct{
	hp int
}

// Method of embedded struct
func (e Engine) Start(){
  fmt.Println("Engine starts with", e.hp, "HP")
}

// Another base struct
type Wheels struct{
	count int
}

// Main struct with EMBEDDED structs
type Car struct {
	Engine   // embedded struct
	Wheels   // embedded struct
	brand string
}

func embeddedStructs() {
	// Creating Car object
	car := Car{
		Engine: Engine{
			hp: 150,
		},
		Wheels: Wheels{
			count: 4,
		},
		brand: "Tesla",
	}

	// Accessing embedded fields directly (FIELD PROMOTION)
	fmt.Println("Car brand:", car.brand)
	fmt.Println("Horsepower:", car.hp)
	fmt.Println("Wheels:", car.count)

	// Calling embedded method directly (METHOD PROMOTION)
	car.Start()
}
