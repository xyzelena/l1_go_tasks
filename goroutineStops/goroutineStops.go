package goroutineStops

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func GoroutineStops() {
	const secToSleep = 3

	fmt.Println("=== Способы остановки выполнения горутины ===")
	fmt.Println()

	fmt.Println("-------- Пример № 1: Остановка через context --------")
	demonstrateContextStop(secToSleep)

	fmt.Println("\n-------- Пример № 2: Остановка через канал уведомления --------")
	demonstrateChannelStop(secToSleep)

	fmt.Println("\n-------- Пример № 3: Остановка по условию в цикле --------")
	demonstrateConditionStop()

	fmt.Println("\n-------- Пример № 4: Остановка через runtime.Goexit() --------")
	demonstrateRuntimeGoexitStop()

	fmt.Println("\n-------- Пример № 5: Остановка через timeout context --------")
	demonstrateTimeoutContextStop()

	fmt.Println("\n-------- Пример № 6: Остановка через atomic флаг --------")
	demonstrateAtomicFlagStop(secToSleep)

	fmt.Println("\n-------- Пример № 7: Остановка через закрытие канала --------")
	demonstrateCloseChannelStop(secToSleep)

	fmt.Println("\n-------- Пример № 8: Остановка через panic/recover --------")
	demonstratePanicRecoverStop()

	fmt.Println("\n=== Демонстрация способов завершена ===")
}

func demonstrateContextStop(secToSleep int) {
	num := atomic.Int32{}
	fmt.Println("Начальное значение num: ", num.Load())

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go increaseNumWithContext(&num, ctx, wg)

	time.Sleep(time.Duration(secToSleep) * time.Second)
	fmt.Println("Отправка сигнала отмены через context...")
	cancel()

	wg.Wait()
	fmt.Println("Финальное значение num: ", num.Load())
}

func demonstrateChannelStop(secToSleep int) {
	num := atomic.Int32{}
	fmt.Println("Начальное значение num: ", num.Load())

	stop := make(chan bool, 1)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go increaseNumWithStop(&num, stop, wg)

	time.Sleep(time.Duration(secToSleep) * time.Second)
	fmt.Println("Отправка сигнала остановки через канал...")
	stop <- true

	wg.Wait()
	close(stop)
	fmt.Println("Финальное значение num: ", num.Load())
}

func demonstrateConditionStop() {
	num := atomic.Int32{}
	const numToStop = 5
	fmt.Println("Начальное значение num: ", num.Load(), " остановка при достижении ", numToStop)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go increaseNumWithCondition(&num, numToStop, wg)
	wg.Wait()

	fmt.Println("Финальное значение num: ", num.Load())
}

func demonstrateRuntimeGoexitStop() {
	num := atomic.Int32{}
	fmt.Println("Начальное значение num: ", num.Load())

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go increaseNumWithRuntimeGoexit(&num, wg)
	wg.Wait()

	fmt.Println("Финальное значение num: ", num.Load(), " (горутина остановлена через runtime.Goexit)")
}

func demonstrateTimeoutContextStop() {
	num := atomic.Int32{}
	fmt.Println("Начальное значение num: ", num.Load())

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go increaseNumWithTimeoutContext(&num, ctx, wg)
	wg.Wait()

	fmt.Println("Финальное значение num: ", num.Load(), " (остановлено по таймауту)")
}

func demonstrateAtomicFlagStop(secToSleep int) {
	num := atomic.Int32{}
	stopFlag := atomic.Bool{}
	fmt.Println("Начальное значение num: ", num.Load())

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go increaseNumWithAtomicFlag(&num, &stopFlag, wg)

	time.Sleep(time.Duration(secToSleep) * time.Second)
	fmt.Println("Установка atomic флага остановки...")
	stopFlag.Store(true)

	wg.Wait()
	fmt.Println("Финальное значение num: ", num.Load())
}

func demonstrateCloseChannelStop(secToSleep int) {
	num := atomic.Int32{}
	fmt.Printf("Начальное значение num: %d\n", num.Load())

	stopCh := make(chan struct{})
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go increaseNumWithCloseChannel(&num, stopCh, wg)

	time.Sleep(time.Duration(secToSleep) * time.Second)
	fmt.Println("Закрытие канала для остановки горутины...")
	close(stopCh)

	wg.Wait()
	fmt.Println("Финальное значение num: ", num.Load())
}

func demonstratePanicRecoverStop() {
	num := atomic.Int32{}
	fmt.Println("Начальное значение num: ", num.Load())

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go increaseNumWithPanicRecover(&num, wg)
	wg.Wait()

	fmt.Println("Финальное значение num: ", num.Load(), " (остановлено через panic/recover)")
}
