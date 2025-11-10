package main

import "fmt"

type person struct {
	name    string
	surname string
	age     int
}

func (p person) Talk() {
	fmt.Println("Привет меня зовут", p.name)
}

func (p *person) incAge() {
	p.age += 1
}

func main() {
	kolya := person{
		name:    "Коля",
		surname: "Пупкин",
		age:     20,
	}

	kolya.Talk()

	fmt.Println("Мой возраст", kolya.age)

	kolya.incAge()

	fmt.Println("А теперь мой возраст", kolya.age)
}
