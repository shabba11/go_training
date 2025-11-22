package main

// Для чего нужен context:
// 1. Хранить значения
// 2. Сообщать о завершении

import (
	"context"
	"fmt"
	"time"
)

func main() {
	BaseKnowledge()
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
