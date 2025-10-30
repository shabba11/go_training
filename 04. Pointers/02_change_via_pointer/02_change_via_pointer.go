package main

import "fmt"

func main() {
	var coffeePrice = 4.50
	fmt.Println("Coffee price", coffeePrice)

	// STEP 1
	// Compile time(code you write):   var coffeePrice = 4.50
	// Runtime (what machine sees):    [some memory address] = 4.50

	// STEP 2
	// Compile time(code you write):   fmt.Println("Coffee price", coffeePrice)
	// Runtime (what machine sees):    fmt.Println([some mem address], [memory address coffeePrice])
	fmt.Println("Memory address of price 4.50", &coffeePrice)

	coffeePrice = 5.00
	fmt.Println("Coffee price: ", coffeePrice)
	fmt.Println("Memory address of price 5.00", &coffeePrice) // same address

	var pointerToCoffeePrice *float64 = &coffeePrice

	/*go to memory location where pointerToCoffeePrice points to
	and change value in this memory location*/
	*pointerToCoffeePrice = 5.50

	fmt.Println("Coffee price: ", *pointerToCoffeePrice)
	fmt.Println("Memory address of price 5.50", pointerToCoffeePrice)
}
