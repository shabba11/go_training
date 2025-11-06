package main

import "fmt"

func greet(coffeeShopName string) {
	var greetMessage string = "Welcome to the Coffe Shop"
	fmt.Println(greetMessage, coffeeShopName)
}

func add(x, y int, a, b, c float32) {
	var z = x + y
	var d = a + b + c
	fmt.Println("x + y = ", z)
	fmt.Println("a + b + c = ", d)
}

func main() {
	greet("Brew & Beans")
	greet("Coffee & Milk")
	add(1, 2, 3.4, 5.6, 1.2)
}
