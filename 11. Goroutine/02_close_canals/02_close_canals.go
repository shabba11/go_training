package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go say("Hello Go!", ch)

	for a := range ch {
		fmt.Println(a)
	}
}
func say(greet string, ch chan int) {
	for i := 0; i <= 5; i++ {
		ch <- i
	}

	close(ch)
}
