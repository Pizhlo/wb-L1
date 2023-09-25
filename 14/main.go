package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

func main() {
	var i int
	fmt.Println(getType(i))

	var s string
	fmt.Println(getType(s))

	var b bool
	fmt.Println(getType(b))

	var ch chan int
	fmt.Println(getType(ch))
	fmt.Println("тип канала: ", reflect.ValueOf(ch).Type())
}

func getType(val interface{}) reflect.Kind {
	return reflect.ValueOf(val).Kind()
}
