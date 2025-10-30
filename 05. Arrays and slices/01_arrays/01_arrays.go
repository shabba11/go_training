package main

import "fmt"

func main() {
	var coffeeSizes [3]string
	fmt.Println(coffeeSizes)

	coffeeSizes[0] = "Small"
	fmt.Println(coffeeSizes)

	coffeeSizes[1] = "Medium"
	fmt.Println(coffeeSizes)

	coffeeSizes[2] = "Large"
	fmt.Println(coffeeSizes)

	coffeeSizes[2] = "Extra Large"
	fmt.Println(coffeeSizes)
	fmt.Println("First element is: ", coffeeSizes[0])

	// coffeeSizes[1] = true // compilation error
	// coffeeSizes[5] = "Super Extra Large" // invalid argument: index 5 out of bounds [0:3]
}
