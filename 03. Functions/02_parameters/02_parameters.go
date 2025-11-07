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

func addSecond(numbers ...int) {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("sum = ", sum)
}

func main() {
	greet("Brew & Beans")
	greet("Coffee & Milk")
	add(1, 2, 3.4, 5.6, 1.2)

	// Неопределенное количество параметров
	// Для определения параметра,
	// который представляет неопределенное количество значений, перед типом этих значений ставится многоточие: numbers ...int
	addSecond(1, 2, 3)       // sum = 6
	addSecond(1, 2, 3, 4)    // sum = 10
	addSecond(5, 6, 7, 2, 3) // sum = 23
	addSecond([]int{1, 2, 3, 4}...)
}
