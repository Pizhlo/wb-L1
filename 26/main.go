package main

import "fmt"

// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// 	aabcd — false

func main() {
	s := "aA"

	fmt.Println(unique(s))
}

func unique(s string) bool {
	m := map[string]struct{}{}

	for i := 0; i < len(s); i++ {
		if _, ok := m[string(s[i])]; ok {
			return false
		}
		m[string(s[i])] = struct{}{}
	}

	return true
}
