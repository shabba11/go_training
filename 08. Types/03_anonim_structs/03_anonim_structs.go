package main

import "fmt"

type person struct {
	string
	// string      // Ошибка: string redeclared - other declaration of string
	int
}

type secondPerson struct {
	string
	company string
	int
}

func main() {

	tom := person{"Tom", 41}
	fmt.Println(tom) // {Tom 41}

	alice := secondPerson{"Alice", "Coca-Cola", 42}
	fmt.Println(alice) // {Tom Coca-Cola 41}

}
