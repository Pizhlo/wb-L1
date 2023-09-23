package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	a := 12
	b := 13

	a, b = b, a

	fmt.Println(a, b)
}
