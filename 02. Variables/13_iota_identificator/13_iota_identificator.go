package main

import "fmt"

func main() {
	const // iota сбрасывается в 0
	(
		C0 = iota // здесь iota равно 0, увеличивается с каждой строкой
		C1        // увеличение на 1, iota равна 1
		C2 = iota // iota равна 2
		C3        // iota равна 2
		C4
		C5 = iota // iota равна 2
	)

	fmt.Println("C0:", C0) // C0: 0
	fmt.Println("C1:", C1) // C1: 1
	fmt.Println("C2:", C2) // C2: 2
	fmt.Println("C3:", C3) // C3: 3
	fmt.Println("C4:", C4) // C4: 4
	fmt.Println("C5:", C5) // C5: 5

	const //  iota сбрасывается в 0
	(
		C6 = iota // С6 = 0
	)

	fmt.Println("C6:", C6) // C6: 0
}
