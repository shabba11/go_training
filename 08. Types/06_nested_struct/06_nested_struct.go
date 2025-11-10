package main

import "fmt"

type person struct { // вложенная структура person
	name string
	age  int
}

type account struct {
	login    string
	password string

	person_info person
}

func main() {

	tom := account{
		login:    "tom@hmail.com",
		password: "12345678",
		person_info: person{
			name: "Tom",
			age:  41,
		},
	}

	fmt.Println(tom) // {tom@hmail.com 12345678 {Tom 41}}

	// обращение к полям вложенной структуры
	fmt.Println("Name: ", tom.person_info.name) // Name: Tom
}
