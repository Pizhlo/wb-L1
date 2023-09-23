package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]bool)

	for _, word := range sequence {
		set[word] = true
	}

	fmt.Println(set)
}
