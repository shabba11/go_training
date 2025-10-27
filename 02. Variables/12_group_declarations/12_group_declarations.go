package main

import "fmt"

func main() {
	var coffeeType = "Latte"
	var quantity = 3
	var unitPrice float64 = 4.25

	fmt.Printf("Ordered %d %s priced $%.2f each\n", quantity, coffeeType, unitPrice)

	var (
		customName   string = "Michail"
		tableNumber  int    = 8
		isReadyToPay bool   = false
	)

	fmt.Printf("Customer: %s with table number: %d. Is ready to pay: %t\n", customName, tableNumber, isReadyToPay)

	// No unused variables compilation error for const
	const (
		sizeSmall  = "S"
		sizeMedium = "M"
		sizeLarge  = "L"
	)
}
