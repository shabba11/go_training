package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	tom := new(person)

	fmt.Println(*tom) // { 0}

	tom.name = "Tom"
	tom.age = 41
	fmt.Println(*tom) // {Tom 41}

	// анонимная структура с new
	bob := new(struct {
		name, company string
		age           int
	})

	fmt.Println(*bob) // {  0}

	bob.name = "Bob"
	bob.company = "Google"
	bob.age = 46

	fmt.Println(*bob) // {Bob Google 46}
}
