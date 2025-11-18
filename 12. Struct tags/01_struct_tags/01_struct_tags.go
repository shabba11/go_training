package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string `validate:"min=2,max=32"`
	Email string `validate:"required,email"` // структурные тэги всегда определяются в обратных ковычках. А значение в двойных ковычках
}

func main() {
	user := User{
		Name:  "Florian",
		Email: "abc@abc.ru",
	}

	t := reflect.TypeOf(user) // Возвращает некоторую структуру отражающую информацию о переменной user
	fmt.Println(t.Name())     // имя
	fmt.Println(t.Kind())     // тип

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag := field.Tag.Get("validate")

		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1,
			field.Name, field.Type.Name(), tag)
	}

}
