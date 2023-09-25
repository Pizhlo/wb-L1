package main

import (
	"fmt"
	"os"
	"time"
)

// Реализовать собственную функцию sleep.

// sleep принимает продолжительность времени, на которое нужно приостановить программу, и приостанавливает выполнение программы,
// пока не получит сообщение в канале
func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Начало программы")

	fmt.Println("Введите количество секунд, на которое приостановить программу:")

	var d time.Duration

	fmt.Fscan(os.Stdin, &d)

	fmt.Printf("Программа останавливается на %d секунд(ы)\n", d)

	sleep(d * time.Second)

	fmt.Printf("Прошло %d секунд(ы)\n", d)
}
