package main

import "fmt"

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var less []int
	var greater []int

	for _, num := range arr[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}

	less = quicksort(less)
	greater = quicksort(greater)

	return append(append(less, pivot), greater...)
}

func main() {
	arr := []int{5, 3, 8, 6, 2, 7, 1, 4}
	fmt.Println(quicksort(arr))
}
