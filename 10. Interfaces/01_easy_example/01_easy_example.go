package main

import "fmt"

type Vehicle interface {
	move()
}

func drive(vehicle Vehicle) {
	vehicle.move()
}

type Car struct{}
type Aircraft struct{}
type Boat struct{}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}
func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}
func (b Boat) move() {
	fmt.Println("Лодка плывет")
}

func main() {

	tesla := Car{}
	boing := Aircraft{}
	jet := Boat{}
	drive(tesla)
	drive(boing)
	drive(jet)
}
