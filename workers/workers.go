package workers

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Worker(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range ch {
		fmt.Println("Worker", id, ":", data)
	}

	fmt.Println("Worker", id, "finished")
}

func DoWorkers(countWorkers int) {
	ch := make(chan int, 10)

	// Контекст для сигналов
	sigCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Таймаут 10 секунд поверх сигнального контекста
	ctx, cancel := context.WithTimeout(sigCtx, 10*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}

	wg.Add(countWorkers)

	for i := 1; i <= countWorkers; i++ {
		go Worker(i, ch, &wg)
	}

	//главный воркер, который генерирует числа
	go func() {
		defer close(ch)

		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	<-ctx.Done()

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Stopped by timeout")
	} else {
		fmt.Println("Stopped by signal")
	}

	wg.Wait()

	fmt.Println("All workers are finished")
}
