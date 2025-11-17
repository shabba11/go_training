package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Указываетя максимальное количество горутин
	runtime.GOMAXPROCS(3)

	// Показывает количество ядер процессора/потоки.
	// То есть максимальное кол-во горутин на данном процессоре

	fmt.Println(runtime.NumCPU())

	go showNumbers(100)

	// Переключение на другую горутину. В данном случае с main на showNumbers
	runtime.Gosched()

	fmt.Println("Exit")
}

func showNumbers(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}
