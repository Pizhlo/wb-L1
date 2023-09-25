package main

import "fmt"

// Удалить i-ый элемент из слайса.

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(delete(arr, 5))
}

func delete(arr []int, i int) []int {
	result := []int{}

	result = append(result, arr[:i-1]...)
	result = append(result, arr[i:]...)

	return result
}
