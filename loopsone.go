package main

import (
	"fmt"
)

func bulkSend(numMessages int) float64 {
    totalCost := 0.0
   for i := 0; i <numMessages; i ++{
	totalCost += 1.0 + (0.01 * float64(i))
   }
   return totalCost
}

// don't edit below this line

func tested(numMessages int) {
	fmt.Printf("Sending %v messages\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulk send complete! Cost = %.2f\n", cost)
	fmt.Println("===============================================================")
}

func loopsone () {
	tested(10)
	tested(20)
	tested(30)
	tested(40)
	tested(50)
}