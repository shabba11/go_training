package main

import (
	"fmt"
	"time"
)

func main() {
	unbufferedChannel := make(chan int) // Такие каналы называются не буферизированные
	// unbufferedChannel := make(chan<- int) // Направленный канал. Только для записи
	// unbufferedChannel := make(<-chan int) // Направленный канал. Только для чтения

	// В небуферизированном канале буфер будет равен = 0. Горутины на запись и чтение будут пустые
	// unbufferedChannel {
	// len(buffer) = 0
	// readers []Goroutines
	// writers []Goroutines
	// }

	fmt.Println(len(unbufferedChannel), cap(unbufferedChannel))

	// unbufferedChannel {
	// len(buffer) = 0
	// readers []Goroutines
	// writers []Goroutines 1
	// }
	// здесь будет deadlock. Для того чтобы использовать небуферизированный канал ->
	// Обязательным условием для записи и чтения из этого канала является
	// то что у нас есть сразу и читатель(readers) и писатель (writers).
	// На каждую записывающую горутину должна быть читающая горутина и наоборот

	// unbufferedChannel <- 3

	// unbufferedChannel {
	// len(buffer) = 0
	// readers []Goroutines 1
	// writers []Goroutines
	// }
	// здесь тоже будет deadlock. Описание почему выше

	// <-unbufferedChannel

	go func(chanForWriting chan<- int) { // здесь указано что chan<- доступен только для записи. <-chan был бы только для чтения
		time.Sleep(time.Second)

		unbufferedChannel <- 1
	}(unbufferedChannel)

	value := <-unbufferedChannel
	fmt.Println(value)

	go func(chanForWriting chan<- int) {
		time.Sleep(time.Second)
		value := <-unbufferedChannel
		fmt.Println(value)
	}(unbufferedChannel)

	unbufferedChannel <- 2

}
