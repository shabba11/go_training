package main

import (
	"fmt"
	"reflect"
)

func main() {
	money := map[string]int{
		"dollars": 1000,
		"euros":   500,
	}

	// var money map[string]int = ....

	fmt.Println(money)
	fmt.Println(money["euros"])

	money["rubles"] = 1400
	fmt.Println(money)

	delete(money, "dollars")
	fmt.Println(money)

	el, status := money["rubles"]
	fmt.Println(el, status)

	el, status = money["kk"]
	fmt.Println(el, status)

	//Функция make представляет альтернативный вариант создания карты. Она создает пустую хеш-таблицу:
	var people = make(map[string]int)
	fmt.Println(people)

	//Для сравнения двух карт по элементам Go предоставляет функцию DeepEqual в пакете reflect. Она возвращает true,
	// если обе карты имеют одинаковые ключи и одинаковые связанные значения ключей, в противном случае возвращается false:
	people1 := map[string]int{"Tom": 1, "Bob": 2, "Sam": 3}
	people2 := map[string]int{"Tom": 1, "Bob": 2, "Sam": 3}
	people3 := map[string]int{"Tom": 1, "Bob": 2, "Sam": 4}
	fmt.Println("people1 == people2:", reflect.DeepEqual(people1, people2)) // people1 == people2: true
	fmt.Println("people1 == people3:", reflect.DeepEqual(people1, people3)) // people1 == people3: false
}
