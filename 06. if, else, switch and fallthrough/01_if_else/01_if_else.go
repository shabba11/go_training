package main

import "fmt"

func main() {

	a := 6
	b := 7
	if a < b {
		fmt.Println("a is less than b")
	}

	//________________________________________//

	if a < b {
		fmt.Println("a меньше b")
	} else {
		fmt.Println("a больше b")
	}

	//________________________________________//

	a = 8
	b = 8
	if a < b {
		fmt.Println("a меньше b")
	} else if a > b {
		fmt.Println("a больше b")
	} else {
		fmt.Println("a равно b")
	}

	//_________________________________________//

	if a == 9 {
		fmt.Println("a = 9")
	} else if a == 8 {
		fmt.Println("a = 8")
	} else if a == 7 {
		fmt.Println("a = 7")
	}
}
