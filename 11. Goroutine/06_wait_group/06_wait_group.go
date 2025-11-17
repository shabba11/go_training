package main

import (
	"fmt"
	"sync"
)

func main() {
	withoutWait()
	withWait()
	wrongAdd()
}

func withoutWait() {
	// Может не успеть вывести все 10 чисел. Из за того что происходит выход из функции
	for i := 0; i < 10; i++ {
		go fmt.Println(i + 1)
	}

	fmt.Println("Exit")
}

func withWait() {
	var wg sync.WaitGroup // Создаем переменную нашей waitGroup'ы

	for i := 0; i < 10; i++ {
		wg.Add(1) // Указываем количество ожидаемых заданий

		go func(i int) {
			fmt.Println(i + 1)
			wg.Done() // Функция сообщает нашей waitGroup'е о том что она завершает свою задачу
		}(i)

	}

	wg.Wait() // Наша горутина блокируется и ждет пока будут выполнены все задачи
	fmt.Println("Exit")
}

func wrongAdd() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {

		go func(i int) {
			wg.Add(1) // внутри горутины не нужно использовать добавление waitGroup'ы. Так как она может не успеть выполниться
			defer wg.Done()
			fmt.Println(i + 1)
		}(i)

	}

	wg.Wait()
}
