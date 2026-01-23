package main

import (
	"fmt"
)

func maxMessages(thresh float64) int {
		totalCost := 0.0
	    count := 0

	for {
		nextCost := 1.0 + float64(count)*0.01

		if totalCost+nextCost > thresh {
			break
		}

		totalCost += nextCost
		count++
	}
	return  count
}

// don't edit below this line

func nest(thresh float64) {
	fmt.Printf("Threshold: %.2f\n", thresh)
	max := maxMessages(thresh)
	fmt.Printf("Maximum messages that can be sent: = %v\n", max)
	fmt.Println("===============================================================")
}

func omit() {
	nest(10.00)
	nest(20.00)
	nest(30.00)
	nest(40.00)
	nest(50.00)
}