package main

import "fmt"

func main() {
	coffeeTypes := [3]string{"Espresso", "Lattee", "Cappucino"}
	fmt.Println("Types of coffee:", coffeeTypes)
	fmt.Println("Lenght of the array:", len(coffeeTypes))

	coffeeTypes[len(coffeeTypes)-1] = "Milk"
	fmt.Println("Types of coffee:", coffeeTypes)
}
