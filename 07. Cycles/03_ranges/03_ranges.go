package main

import "fmt"

func main() {

	// for индекс, значение := range набор_данных{
	// 		действия
	// }

	nameHello := "Hello"
	for index, value := range nameHello {
		fmt.Println("Index:", index, "Value:", value)
	}

	// Index: 0  Value: 72
	// Index: 1  Value: 101
	// Index: 2  Value: 108
	// Index: 3  Value: 108
	// Index: 4  Value: 111
	// Здесь мы видим, что значением каждого символа строки является его числовой код.
	// Но попробуем перевести код символа в читабельную форму

	for index, value := range nameHello {
		fmt.Printf("Index: %d, Value: %c\n", index, value)
	}

	// Index: 0, Value: H
	// Index: 1, Value: e
	// Index: 2, Value: l
	// Index: 3, Value: l
	// Index: 4, Value: o

	// Если мы не планируем использовать значения или индексы элементов,
	// то мы можем вместо них указать прочерк. Например, нам не нужны индексы

	for _, value := range nameHello {
		fmt.Printf("%c ", value)
	}
	fmt.Printf("%c ", 10)

	// Аналогично мы можем перебирать и остальные типы наборов данных. Например, перебирем массив строк.
	var users = [3]string{"Tom", "Alice", "Kate"}
	for index, value := range users {
		fmt.Println(index, value)
	}

	// 0 Tom
	// 1 Alice
	// 2 Kate

	// Для перебора массива можно использовать и стандартную версию цикла for

	var usersSecond = [3]string{"Tom", "Alice", "Kate"}
	for i := 0; i < len(usersSecond); i++ {
		fmt.Println(usersSecond[i])
	}
}
