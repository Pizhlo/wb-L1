package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	arr := []int{2, 4, 6, 8, 10}

	calculateSquare(arr)
}

func calculateSquare(arr []int) {
	var wg sync.WaitGroup

	for _, val := range arr {
		wg.Add(1)
		go func(item int) {
			fmt.Println(item * item)
			defer wg.Done()
		}(val)

	}
	wg.Wait()
}
