package timeoutChannel

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func TimeoutChannel() {
	ch := make(chan int, 10)

	const N = 3

	ctx, cancel := context.WithTimeout(context.Background(), N*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}

	// sender
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)

		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("timeout after %d seconds — shutting down\n", N)
				fmt.Println("sender: stopping")
				return
			case ch <- i:
				fmt.Println("send value -->", i)
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	//receiver
	for v := range ch {
		fmt.Println("--> receive value:", v)
	}

	wg.Wait()
}
