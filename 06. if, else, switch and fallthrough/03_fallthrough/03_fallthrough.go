package main

import "fmt"

func main() {
	/*
		По умолчанию Go выходит из switch после первого совпадения.
		Но оператор fallthrough в конструкции switch заставляет выполнение переходить от одного случая к другому,
		даже если последующие случаи не совпадают
	*/
	x := 2

	switch x {
	case 1:
		fmt.Println("x = 1")
	case 2:
		fmt.Println("x = 2")
		fallthrough
	case 3:
		fmt.Println("x = 3")
	case 4:
		fmt.Println("x = 4")
	}

	// x = 2
	// x = 3
}
