package main

import "fmt"

func main() {

	a := 8
	switch a {
	case 9:
		fmt.Println("a = 9")
	case 8:
		fmt.Println("a = 8")
	case 7:
		fmt.Println("a = 7")
	}

	//____________________________________________//

	a = 7
	switch a + 2 {
	case 9:
		fmt.Println("9")
	case 8:
		fmt.Println("8")
	case 7:
		fmt.Println("7")
	}

	//______________________________________________//

	a = 87
	switch a {
	case 9:
		fmt.Println("a = 9")
	case 8:
		fmt.Println("a = 8")
	case 7:
		fmt.Println("a = 7")
	default:
		fmt.Println("значение переменной a не определено")
	}

	//______________________________________________//

	a = 5
	switch a {
	case 9:
		fmt.Println("a = 9")
	case 7:
		fmt.Println("a = 7")
	case 6, 5, 4:
		fmt.Println("a = 6 or a = 5 or a = 4")
	default:
		fmt.Println("значение переменной a не определено")
	}
}
