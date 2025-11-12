package main

import (
	"fmt"
)

func main() {
	ch := make(chan string) // Создаем канал
	ch2 := make(chan int)

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
