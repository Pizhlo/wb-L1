package main

import "fmt"

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func main() {
	stringArr := []string{"snow", "dog", "sun"}

	fmt.Println(reverse(stringArr))
}

func reverse(arr []string) []string {
	res := []string{}

	for i := len(arr) - 1; i >= 0; i-- {
		res = append(res, arr[i])
	}
	return res
}
