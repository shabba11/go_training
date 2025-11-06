package main

import "fmt"

func main() {

	// Одномерные массивы
	coffeeTypes := [3]string{"Espresso", "Lattee", "Cappucino"}
	fmt.Println("Types of coffee:", coffeeTypes)
	fmt.Println("Lenght of the array:", len(coffeeTypes))

	coffeeTypes[len(coffeeTypes)-1] = "Milk"
	fmt.Println("Types of coffee:", coffeeTypes)

	anotherCoffeeTypes := [...]string{"Lattee", "Cappucino"}
	fmt.Println("Types of coffee:", anotherCoffeeTypes)
	fmt.Println("Lenght of the array:", len(anotherCoffeeTypes))

	colors := [3]string{2: "blue", 0: "red", 1: "green"}
	fmt.Println(colors[2]) // blue

	// Многомерные массивы
	numbers := [3][2]int{
		{1, 2},
		{4, 5},
		{7, 8},
	}
	fmt.Println(numbers) // [[1 2] [4 5] [7 8]]

	//___________________________________________//

	numbersSecond := [3][2]int{
		{1, 2},
		{5},
	}

	fmt.Println(numbersSecond) // [[1 2] [5 0] [0 0]]

	//___________________________________________//

	var numbersThird [3][2]int

	numbersThird[0] = [2]int{1, 2}
	numbersThird[1] = [2]int{4, 5}
	numbersThird[2] = [2]int{7, 8}

	fmt.Println(numbersThird) // [[1 2] [4 5] [7 8]]

	//___________________________________________//

	var numbersFour [3][2]int

	numbersFour[0][0] = 1
	numbersFour[0][1] = 2

	numbersFour[1][0] = 4
	numbersFour[1][1] = 5
	numbersFour[2][0] = 7
	numbersFour[2][1] = 8

	fmt.Println(numbersFour) // [[1 2] [4 5] [7 8]]
}
