package main

import "fmt"

type person struct {
	name string
	age  int
}

type account struct {
	login    string
	password string
	person
}

func main() {

	tom := account{
		"tom@hmail.com",
		"12345678",
		person{"Tom", 41},
	}

	fmt.Println(tom) // {tom@hmail.com 12345678 {Tom 41}}

	// обращение к полям встроенной структуры
	fmt.Println("Name: ", tom.person.name) // Name: Tom

}
