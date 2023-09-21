package main

import "fmt"

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	var num int64
	var bitIndex int
	var bitValue int

	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	fmt.Print("Введите индекс бита (0-63): ")
	fmt.Scan(&bitIndex)

	fmt.Print("Введите значение, которое нужно установить (0 or 1): ")
	fmt.Scan(&bitValue)

	if bitValue == 1 {
		num |= 1 << bitIndex
	} else if bitValue == 0 {
		num &= ^(1 << bitIndex)
	}

	fmt.Printf("Результат: %d\n", num)
}
