package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Реализовать конкурентную запись данных в map.

func main() {
	var m sync.Map

	go write(&m)
	go read(&m)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

}

func read(m *sync.Map) {
	for i := 0; ; i++ {
		val, ok := m.Load(i)
		if ok {
			fmt.Println("Загружено значение: ", val)
		}
	}
}

func write(m *sync.Map) {
	for i := 0; ; i++ {
		m.Store(i, i)
	}
}
