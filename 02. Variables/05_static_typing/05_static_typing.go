package main

import "fmt"

func main() {
	// Explicit type declaration
	var cupsQty int = 3

	// capsQty = "four" // Compile-time error
	fmt.Println("Number of cups:", cupsQty)

	// Implicit type declaration
	var wasProcessed = true

	// wasProcessed = "yes" // Compile-time error
	fmt.Println("Order was processed:", wasProcessed)
}
