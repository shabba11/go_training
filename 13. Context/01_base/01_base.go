package main

// Для чего нужен context:
// 1. Хранить значения
// 2. Сообщать о завершении

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	BaseKnowledge()
	workerPool()
}

func BaseKnowledge() {
	ctx := context.Background() // Background создает корневой контекст - самый основной
	fmt.Println(ctx)

	toDo := context.TODO() // Считается что TODO используется в тестах.
	// Или в те моменты когда мы не знаем пригодится нам что нибудь или нет
	fmt.Println(toDo)

	withValue := context.WithValue(ctx, "name", "vasya")
	fmt.Println(withValue.Value("name"))

	withCancel, cancel := context.WithCancel(ctx) // cancel - это функция которая умеет завершать этот контекст
	fmt.Println(withCancel.Err())
	cancel()
	fmt.Println(withCancel.Err())

	withDeadline, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3)) // WithDeadline - можно установить дедлайн контекста
	defer cancel()
	fmt.Println(withDeadline.Deadline())
	fmt.Println(withDeadline.Err())
	fmt.Println(<-withDeadline.Done())

	withTimeOut, cancel := context.WithTimeout(ctx, time.Second*3) // Тоже самое но не со временем завершения а через сколько секунд завершить.
	defer cancel()
	fmt.Println(<-withTimeOut.Done())
}

func workerPool() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup()
	numbersToProcess, processedNumbers := make(chan int, 5), make(chan int, 5)
	// numbersToProcess - в этом канале будет лежать числа которые нужно обработать,
	// processedNumbers - числа которые мы уже обработали

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, numbersToProcess, processedNumbers)
		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			if i == 500 {
				cancel()
			}
			numbersToProcess <- i
		}
		close(numbersToProcess)
	}()

	go func() {
		wg.Wait()
		close(processedNumbers)
	}()

	var counter int
	for resultValue := range processedNumbers {
		counter++
		fmt.Println(resultValue)
	}
}

func worker(ctx context.Context, toProcess <-chan int, processed chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case value, ok := <-toProcess:
			if !ok {
				return
			}
			time.Sleep(time.Millisecond)
			processed <- value * value
		}
	}
}
