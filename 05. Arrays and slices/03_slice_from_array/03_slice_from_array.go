package main

import "fmt"

func main() {
	dessertMenu := [4]string{"Muffin", "Brownie", "Croissant", "Cookie"}
	fmt.Println("Desset menu:", dessertMenu)

	slice := dessertMenu[1:3] // elements with indexes 1 and 2
	fmt.Println("Slice of the dessert menu [1:3]:", slice)

	slice = dessertMenu[:] // all elements
	fmt.Println("Slice of the dessert menu [:]:", slice)

	slice = dessertMenu[2:] // all elements starting from index 2
	fmt.Println("Slice of the dessert menu [2:]:", slice)

	slice = dessertMenu[:3] // all elements from the beginning to index 3 (not including)
	fmt.Println("Slice of the dessert menu [:3]:", slice)
}
