package main

import "fmt"

type Numbers struct {
	num1 int
	num2 int
}

type NumberInterface interface {
	Sum() int
	Multiply() int
	Division() float64
	Substract() int
}

func (n Numbers) Sum() int {
	return n.num1 + n.num2
}
func (n Numbers) Multiply() int {
	return n.num1 * n.num2
}
func (n Numbers) Division() float64 {
	return float64(n.num1) / float64(n.num2)
}
func (n Numbers) Substract() int {
	return n.num1 - n.num2
}

func main() {
	var i NumberInterface
	numbers := Numbers{9, 8}
	i = numbers
	fmt.Println(i.Sum())
	fmt.Println(i.Multiply())
	fmt.Println(i.Division())
	fmt.Println(i.Substract())
}
