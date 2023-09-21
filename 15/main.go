package main

import (
	"bytes"
	"fmt"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

// Код в примере приводит к утечке памяти, потому что создается большая строка, но используются только первые 100 элементов.
// Для избежания утечки лучше использовать буферизированный поток вывода (например, bytes.Buffer), чтобы записать первые 100 символов в буфер,
// а затем освободить память, выделенную для большой строки.

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	buf := bytes.NewBufferString(v[:100])
	justString = buf.String()
}

func main() {
	someFunc()
	fmt.Println(justString)
}
