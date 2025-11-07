package main

import "fmt"

func sum(n uint) uint {

	if n == 1 {
		return n
	}
	return n + sum(n-1)
}
func main() {

	fmt.Println(sum(4)) // 10
	fmt.Println(sum(5)) // 15
	fmt.Println(sum(6)) // 21

	// Например, вызов sum(4) фактически можно расписать следующим образом:

	// sum(4)

	// 4 + sum(3)

	// 4 + 3 + sum(2)

	// 4 + 3 + 2 + sum(1)

	// 4 + 3 + 2 + 1
}
