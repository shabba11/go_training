package main

import "fmt"

func main() {
	bufferedChannel := make(chan int, 2)                    // вторым аргументом мы указываем размер буфера
	fmt.Println(len(bufferedChannel), cap(bufferedChannel)) // cap-капатиси канала - это его буфер. len - отображает кол-во элементов в буфере

	// Не будет ошибки так как значения запишуться в буфер
	bufferedChannel <- 1
	bufferedChannel <- 2
	fmt.Println(len(bufferedChannel), cap(bufferedChannel))

	// Будет ошибка, так как в буфере нет места
	// bufferedChannel <- 3

	// Извлекаем значения из канала
	fmt.Println(<-bufferedChannel) // Здесь буфер освобождается и появляется свободное место
	bufferedChannel <- 3
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)

	// Будет ошибка. Так как значений в канале больше нет. Deadlock
	// fmt.Println(<-bufferedChannel)

}
