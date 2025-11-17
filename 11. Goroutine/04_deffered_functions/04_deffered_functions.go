package main

import "fmt"

func main() {
	// Складывается в так называемый defer стек и выполняется после завершения основного потока.
	// Выполняются в обратном порядке (с последнего добавленного в стек)
	defer fmt.Println(1)
	defer fmt.Println(2)

	// На выходе из функции сработает defer и вернет нам 10.
	fmt.Println(sum(2, 3))

	// На выходе из функции так же сначала отобразится второй цикл а позже первый
	deferValues()

	fmt.Println("Exit")
}

func sum(x, y int) (sum int) {
	defer func() {
		sum *= 2
	}()

	sum = x + y
	return
}

func deferValues() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("first", i)
	}
	for i := 0; i < 10; i++ {
		// Будет все время отображаться 10 (так как последняя запись была 10). Так выводить некорректно
		// правильней будет вызвать переменную n := 0 вне цикла и увеличивать ее при каждой итерации
		defer func() {
			fmt.Println("second", i)
		}()

		// Либо вызвать ее таким образом
		// defer func(k int) {
		// 	fmt.Println("fourth", k)
		// }(i) // i будет использован как аргумента для параметра функции
	}
}
