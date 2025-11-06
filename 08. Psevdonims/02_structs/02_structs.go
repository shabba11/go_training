package main

import "fmt"

type person struct {
	name string
	age  int
}

var fred struct {
	name string
	age  int
}

func main() {

	var tom person = person{"Tom", 23}
	fmt.Println(tom) // {Tom 23}

	var alice person = person{age: 23, name: "Alice"}
	fmt.Println(alice) // {Alice 23}

	bob := person{name: "Bob", age: 31}
	fmt.Println(bob) // {Bob 31}

	undefined := person{} // name - пустая строка, age - 0
	fmt.Println(undefined)

	//________________________________________//

	var ivan = person{name: "Ivan", age: 24}
	fmt.Println(ivan.name) // Tom
	fmt.Println(ivan.age)  // 24

	tom.age = 38                     // изменяем значение
	fmt.Println(ivan.name, ivan.age) // Ivan 38

	// Необязательно объявлять новый тип при объявлении структур.
	// Мы также можем создать анонимную структуру в языке Go

	fred.name = "Fred"
	fred.age = 23
	fmt.Println(fred.name) // Fred
	fmt.Println(fred.age)  // 23

	//_________________________________________//

	// Также мы можем сразу инициализировать переменную анонимной структуры

	sasha := struct { // объявление переменной анонимной структуры
		name string
		age  int
	}{ // инициализация переменной
		name: "Sasha",
		age:  41,
	}

	fmt.Println(sasha.name) // Sasha
	fmt.Println(sasha.age)  // 41
}
