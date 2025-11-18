package main

import (
	"fmt"
)

func main() {
	ch := make(chan string) // Создаем канал. Такой канал называется не буферизированный
	ch2 := make(chan int)

	// канал можно создать задав переменную
	var nilChannel chan int
	nilChannel <- 3
	// у нашего канала nilChannel (канал nil'овый) нет ни буфера ни очередей,
	// поэтому переменную некуда поместить и будет ошибка
	// сlose(nilChannel) // nil'овый канал закрыть тоже не получится

	go say("Hello Go!", ch, ch2)
	// time.Sleep(10 * time.Second)
	data1 := <-ch
	data2 := <-ch2
	fmt.Println(data1, data2) // Go ожидает получение данных с канала
}
func say(greet string, ch chan string, ch2 chan int) {
	fmt.Println(greet)
	ch <- "Привет от горутины!"
	ch2 <- 74
}
