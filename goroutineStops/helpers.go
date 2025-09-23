package goroutineStops

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func increaseNumWithContext(num *atomic.Int32, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Получен сигнал отмены через context")
			return
		default:
			num.Add(1)
			fmt.Println("- Context: Num ", num.Load())
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func increaseNumWithStop(num *atomic.Int32, stop chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-stop:
			fmt.Println("Получен сигнал остановки через канал")
			return
		default:
			num.Add(1)
			fmt.Println("- Channel: Num ", num.Load())
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func increaseNumWithCondition(num *atomic.Int32, numToStop int32, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		if num.Load() >= numToStop {
			fmt.Println("- Ограничение на количество итераций достигнуто! (num >= ", numToStop, ")")
			return
		}

		num.Add(1)
		fmt.Println("- Condition: Num ", num.Load())
		time.Sleep(500 * time.Millisecond)
	}
}

func increaseNumWithRuntimeGoexit(num *atomic.Int32, wg *sync.WaitGroup) {
	defer wg.Done()

	num.Add(1)
	fmt.Println("- Runtime.Goexit: Num ", num.Load())
	fmt.Println("Вызов runtime.Goexit() - немедленное завершение горутины")
	runtime.Goexit()
}

func increaseNumWithTimeoutContext(num *atomic.Int32, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Println("Таймаут контекста истек!")
			} else {
				fmt.Println("Контекст отменен!")
			}
			return
		default:
			num.Add(1)
			fmt.Println("- TimeoutContext: Num ", num.Load())
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func increaseNumWithAtomicFlag(num *atomic.Int32, stopFlag *atomic.Bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		if stopFlag.Load() {
			fmt.Println("Atomic флаг остановки установлен!")
			return
		}

		num.Add(1)
		fmt.Println("- AtomicFlag: Num ", num.Load())
		time.Sleep(500 * time.Millisecond)
	}
}

func increaseNumWithCloseChannel(num *atomic.Int32, stopCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-stopCh:
			fmt.Println("Канал закрыт, горутина остановлена!")
			return
		default:
			num.Add(1)
			fmt.Println("- CloseChannel: Num ", num.Load())
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func increaseNumWithPanicRecover(num *atomic.Int32, wg *sync.WaitGroup) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("- Перехвачена паника: ", r)
		}
		wg.Done()
	}()

	for {
		num.Add(1)
		fmt.Println("- PanicRecover: Num ", num.Load())

		if num.Load() >= 3 {
			panic("Преднамеренная паника для остановки горутины!")
		}

		time.Sleep(500 * time.Millisecond)
	}
}
