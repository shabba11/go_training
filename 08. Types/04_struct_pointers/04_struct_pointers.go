package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	tom := person{name: "Tom", age: 22}
	var p_tom *person = &tom // p_tom - указатель на

	// Обращаемся к полям структуры через указатель

	fmt.Println(p_tom.name) // Tom
	fmt.Println(p_tom.age)  // 22

	p_tom.age = 23 // меняем значение поля через указатель
	// (*p_tom).age = 23
	fmt.Println(tom.age) // 23

	// Также можно определять указатели на отдельные поля структуры

	tom = person{name: "Tom", age: 22}
	var p_age *int = &tom.age // указатель на поле tom.age
	*p_age = 35               // изменяем значение поля
	fmt.Println(tom.age)      //  35
}
