package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

func reverse(s string) string {
	var res string

	for _, v := range s {
		res = string(v) + res
	}

	return res
}

func main() {
	s := "главрыба"

	fmt.Println(reverse(s))
}
