package main

import (
	"fmt"
	"sync"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func main() {
	inputChan := make(chan int)
	outputChan := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(2)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	go func() {
		defer wg.Done()
		for _, val := range arr {
			inputChan <- val
		}
		close(inputChan)
	}()

	go func() {
		defer wg.Done()
		for num := range inputChan {
			result := num * 2
			outputChan <- result
		}
		close(outputChan)
	}()

	wg.Wait()

	for res := range outputChan {
		fmt.Println(res)
	}
}
