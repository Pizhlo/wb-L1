package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func main() {
	go boolVar()
}

// 1. Использование булевой переменной:

func boolVar() {
	ch := make(chan int)
	stop := make(chan bool)

	go func() {
		for i := 0; ; i++ {
			select {
			case ch <- i:
			case <-stop:
				return
			}
		}
	}()

	go func() {
		time.Sleep(10 * time.Second)
		stop <- true
	}()
}

// 2. Использование контекста:

func withContext() {
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for i := 0; ; i++ {
			select {
			case ch <- i:
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		time.Sleep(10 * time.Second)
		cancel()
	}()
}

// 3. Использование сигнала ОС:

func osSignal() {
	ch := make(chan int)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	go func() {
		for i := 0; ; i++ {
			select {
			case ch <- i:
			case <-sig:
				return
			}
		}
	}()

	go func() {
		time.Sleep(10 * time.Second)
		sig <- os.Interrupt
	}()
}

// 4. Использование селекта с таймаутом:

func selectWithTimeout() {
	ch := make(chan int)

	go func() {
		for i := 0; ; i++ {
			select {
			case ch <- i:
			case <-time.After(1 * time.Second):
			}
		}
	}()

	go func() {
		time.Sleep(10 * time.Second)
		close(ch)
	}()

}

// 5. Использование буферизованного канала:

func withChannel() {
	ch := make(chan int, 1)

	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	go func() {
		time.Sleep(10 * time.Second)
		close(ch)
	}()
}

// 6. Использование sync.WaitGroup:

func withWG() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; ; i++ {
			select {
			case ch <- i:
			case <-time.After(1 * time.Second):
			}
		}
	}()

	go func() {
		time.Sleep(10 * time.Second)
		close(ch)
	}()

	wg.Wait()
}
