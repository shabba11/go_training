package main

import "fmt"

type Writer interface {
	write(string)
}

type File struct {
	text string
}

// реализация интерфейса Writer для *File
func (f *File) write(message string) {
	f.text = message
	fmt.Println("Запись в файл строки", message)
}

func main() {

	myFile := File{"Undefined"}
	myFile.write("Hello World")
	fmt.Println(myFile.text) // Hello World
}
