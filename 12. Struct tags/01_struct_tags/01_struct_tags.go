package main

// Struct tags — это аннотации в виде строк, привязанные к полям структуры.
// Они задаются после объявления типа поля в обратных кавычках и обычно используются
// для передачи метаданных библиотекам (например, кодировщикам JSON, ORM, валидаторам).

import (
	"fmt"
	"reflect"
)

// Cтандартная библиотека reflect извлекает теги через reflect.Type.Field(i).Tag и методы Tag.Get("key").

// type User struct {
//     ID    int    `json:"id"`
//     Name  string `json:"name,omitempty"`
//     Email string `db:"email" validate:"required,email"`
// }

// t := reflect.TypeOf(User{})
// field, _ := t.FieldByName("Email")
// fmt.Println(field.Tag.Get("db")) // "email"
// fmt.Println(field.Tag.Get("validate")) // "required,email"

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
