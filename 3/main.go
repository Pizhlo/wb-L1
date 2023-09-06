package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

func main() {
	arr := []int{2, 4, 6, 8, 10}

	fmt.Println(calcSum(arr))
}

func calcSum(arr []int) int {
	var wg sync.WaitGroup
	var res int

	for _, val := range arr {
		wg.Add(1)
		go func(item int) {
			res += item * item
			wg.Done()
		}(val)
		wg.Wait()
	}

	return res

}
