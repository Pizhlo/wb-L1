package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают
// произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func main() {
	fmt.Println("Введите количество воркеров:")

	var workerCount int

	fmt.Fscan(os.Stdin, &workerCount)

	ch := make(chan int)

	// запускаем главный поток, который будет записывать данные в канал
	go func() {
		writeData(ch)
	}()

	wg := sync.WaitGroup{}
	wg.Add(workerCount)

	// создаем воркеры
	for i := 0; i < workerCount; i++ {
		go func() {
			defer wg.Done()

			// читаем данные из канала и выводим в stdout
			readData(ch)
		}()
	}

	// ожидаем сигнал завершения программы
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	// закрываем канал и ожидаем завершения всех воркеров
	close(ch)
	wg.Wait()
	fmt.Println("All workers finished")
}

func writeData(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func readData(ch <-chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}
