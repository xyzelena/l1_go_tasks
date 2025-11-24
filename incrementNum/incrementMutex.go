package incrementNum

import (
	"fmt"
	"sync"
)

type CounterMutex struct {
	mu sync.Mutex
	a  int
}

func (c *CounterMutex) doIncrement() {
	c.mu.Lock()
	c.a += 1
	c.mu.Unlock()
}


func (c *CounterMutex) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.a
}

func IncrementMutex() {
	var (
		counter CounterMutex
		wg sync.WaitGroup
	)

	println("Increment with mutex. Initial value: ", counter.getValue())

	wg.Add(100)

	for range 100 {
		go func() {
			defer wg.Done()

			counter.doIncrement()
		}()
	}

	wg.Wait()

	fmt.Println("Increment with mutex: ", counter.a)
}

