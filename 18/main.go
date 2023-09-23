package main

import (
	"fmt"
	"sync"
	"time"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

type Counter struct {
	mux  sync.RWMutex
	data int
}

func New() *Counter {
	return &Counter{
		data: 0,
	}
}

func (c *Counter) Increment() {
	c.mux.Lock()
	c.data++
	c.mux.Unlock()
}

func (c *Counter) GetData() int {
	c.mux.RLock()
	res := c.data
	c.mux.RUnlock()
	return res
}

func main() {
	counter := New()
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ { // запускаем 5 горутин, инкрементирующих счетчик
		go counter.Increment()
		wg.Add(1)
		go func() {
			counter.Increment()
			defer wg.Done()
		}()
	}

	wg.Wait()

	time.Sleep(30 * time.Second)

	fmt.Println(counter.GetData())
}
