package main

import (
	"fmt"
	"os"
	"sort"
)

// Реализовать бинарный поиск встроенными методами языка.

func main() {
	arr := []int{7, 2, 5, 6, 11, 3, 2, 9} // массив, в котором будем проводить поиск
	fmt.Println("Введите число, которое нужно найти:")
	var number int

	fmt.Fscan(os.Stdin, &number)

	sort.Slice(arr, func(i, j int) bool { // отсортируем слайс, чтобы подготовить к бинарному поиску
		return arr[i] < arr[j]
	})

	// 2 2 3 5 6 7 9 11
	fmt.Println(find(arr, number))
}

// find ищет элемент в массиве и возвращает индекс (если элемент не найден - возвращает -1)
func find(arr []int, num int) int {
	minIndex := 0
	maxIndex := len(arr) - 1
	midIndex := 0

	for maxIndex >= minIndex {
		midIndex = (minIndex + maxIndex) / 2
		if arr[midIndex] == num {
			return midIndex
		} else if arr[midIndex] < num {
			minIndex += 1
		} else if arr[midIndex] > num {
			maxIndex -= 1
		}
	}
	return -1
}
