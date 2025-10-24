package pipeline

import (
	"fmt"
)

func Pipeline() {
	var arr = []int{1, 2, 3, 4, 5}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for _, x := range arr {
			ch1 <- x
		}
		close(ch1)
	}()

	go func() {
		for v := range ch1 {
			ch2 <- v * 2
		}
		close(ch2)
	}()

	for v := range ch2 {
		fmt.Println(v)
	}

}
