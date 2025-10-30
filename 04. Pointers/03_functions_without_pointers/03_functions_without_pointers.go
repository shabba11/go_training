package main

import "fmt"

func calculatePriceAfterDiscount(price float64, discountRate float64) float64 {
	return price - (price * discountRate)
}

func main() {
	var coffeePrice float64 = 5.00
	var dicsount float64 = 0.10
	fmt.Printf("Basic coffee price: $%.2f\n", coffeePrice)

	coffeePrice = calculatePriceAfterDiscount(coffeePrice, dicsount)
	fmt.Printf("Price with discount: $%.2f\n", coffeePrice)
}
