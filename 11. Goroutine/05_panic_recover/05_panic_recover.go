package main

import "fmt"

func main() {
	makePanic()
}

func makePanic() {
	defer func() {
		panicValue := recover() // Результат значения panicValue будет значение panic, если panic произошел
		fmt.Println(panicValue)
	}()

	panic("some panic")             // вызов ошибки
	fmt.Println("Unreachable code") // до сюда выполнение функции из за panic просто не дойдет
}
