package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {
	a := big.NewInt(1<<20 + 4)
	b := big.NewInt(1<<20 + 2)

	fmt.Println("a =", a)
	fmt.Println("b =", b)

	division := big.NewInt(0).Div(a, b)
	sum := big.NewInt(0).Add(a, b)
	multiplication := big.NewInt(0).Mul(a, b)
	subtraction := big.NewInt(0).Sub(a, b)

	fmt.Println("division =", division)
	fmt.Println("sum =", sum)
	fmt.Println("multiplication =", multiplication)
	fmt.Println("subtraction =", subtraction)
}
