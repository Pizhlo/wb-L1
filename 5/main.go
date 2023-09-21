package main

import (
	"fmt"
	"os"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func main() {
	ch := make(chan int)

	fmt.Println("Введите количество секунд:")

	var sec time.Duration

	fmt.Fscan(os.Stdin, &sec)

	// запускаем главный поток, который будет записывать данные в канал
	go func() {
		for i := 0; ; i++ {
			ch <- i
			time.Sleep(time.Second) // ждем 1 секунду перед записью следующего значения
		}
	}()

	// запускаем горутину для чтения данных из канала
	go func() {
		for val := range ch {
			fmt.Println(val)
		}
	}()

	// ожидаем N секунд и завершаем программу
	time.Sleep(sec * time.Second)
	fmt.Println("Program finished")
}
