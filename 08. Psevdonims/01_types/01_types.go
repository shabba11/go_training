package main

import "fmt"

type mile uint
type kilometer uint

func distanceToEnemy(distance mile) {

	fmt.Println("расстояние для противника:")
	fmt.Println(distance, "миль")
}

func main() {

	var distance mile = 5
	distanceToEnemy(distance)
	// var distance2 kilometer = 5
	// distanceToEnemy(distance2)   // ! ошибка
}
