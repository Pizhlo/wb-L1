package main

import (
	"fmt"
)

// Реализовать пересечение двух неупорядоченных множеств.

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	intersection := make(map[int]bool)

	for _, num := range set1 {
		intersection[num] = true
	}

	for _, num := range set2 {
		if intersection[num] {
			fmt.Println(num)
		}
	}
}
