package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	writeWithoutConcurrent()
	writeWithoutMutex()
	writeWithMutex()
	readWithMutex()
}

func writeWithoutConcurrent() {
	start := time.Now()
	var counter int

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Nanosecond)
		counter++
	}

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func writeWithoutMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done() // Очень удобно закрывать задачу defer'ом. Чтобы не вызывать ее в конце
			time.Sleep(time.Nanosecond)
			counter++ // здесь будет происходить гонка данных (Data Race).
			// То есть одновременно могут записываться одинаковые данные.
			// В одном потоке было 555+1=556 а во втором было тоже самое и запишеться одно и тоже число
		}()
	}
	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func writeWithMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex // По сути Mutex это эксклюзивная блокировка на запись данных

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func readWithMutex() {
	start := time.Now()
	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.Mutex
	)

	wg.Add(100)

	for i := 0; i < 50; i++ { // Представим что у нас 50 операций чтения
		go func() {
			defer wg.Done()

			mu.Lock()

			time.Sleep(time.Nanosecond)
			_ = counter // Имитация чтения

			mu.Unlock()
		}()
	}

	for i := 0; i < 50; i++ { // Представим что у нас 50 операций записи
		go func() {
			defer wg.Done()

			mu.Lock()

			time.Sleep(time.Nanosecond)
			counter++ // Имитация записи

			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}
