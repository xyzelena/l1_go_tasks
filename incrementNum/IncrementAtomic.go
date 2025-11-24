package incrementNum

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type CounterAtomic struct {
	a atomic.Int32
}

func (c *CounterAtomic) doIncrement() {
	c.a.Add(1)
}

func (c *CounterAtomic) getValue() int {
	return int(c.a.Load())
}

func IncrementAtomic() {
	var (
		counter CounterAtomic
		wg sync.WaitGroup
	)

	println("Increment with atomic. Initial value: ", counter.getValue())

	wg.Add(100)
	
	for range 100 {
		go func() {
			defer wg.Done()
			counter.doIncrement()
		}()
	}

	wg.Wait()

	fmt.Println("Increment with atomic: ", counter.getValue())
}