package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	readWithRWMutex()
}

func readWithRWMutex() {
	start := time.Now()
	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.RWMutex
	)

	wg.Add(100)

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			mu.RLock()

			time.Sleep(time.Nanosecond)
			_ = counter // Тут происходит блокировка именно на чтение

			mu.RUnlock()
		}()
	}

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			mu.Lock()

			time.Sleep(time.Nanosecond)
			counter++ // Тут про

			mu.Unlock()
		}()

	}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}
