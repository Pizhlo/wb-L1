package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

func main() {
	h := NewHuman("name", 1)
	action := NewAction(h)

	fmt.Println(h.Name())
	fmt.Println(action.Name())

	fmt.Println(h.Age())
	fmt.Println(action.Age())

	h.SetAge(5)
	fmt.Println(h.Age())
	fmt.Println(action.Age())

	action.SetAge(10)
	fmt.Println(action.Age())
	fmt.Println(h.Age())

}
