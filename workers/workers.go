package workers

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Worker(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range ch {
		fmt.Println("Worker", id, ":", data)
	}

	fmt.Println("Worker", id, "завершён")
}

func DoWorkers(countWorkers int) {
	ch := make(chan int, 10)

	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}

	wg.Add(countWorkers)

	for i := 1; i <= countWorkers; i++ {
		go Worker(i, ch, &wg)
	}

	// Главная горутина постоянно записывает данные в канал
	go func() {
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(5 * time.Second)

	cancel()
	close(ch)
	wg.Wait()

	fmt.Println("Все воркеры завершены")
}
