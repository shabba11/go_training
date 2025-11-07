package main

import "fmt"

func main() {
	var pf *float64 // nil pointer - то есть пустой
	if pf != nil {
		fmt.Println("Value:", *pf)
	}
}
