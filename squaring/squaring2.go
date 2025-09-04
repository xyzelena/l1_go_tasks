// squaring with wait group

package squaring

import (
	"fmt"
	"sync"
)

func Squaring2() {
	initArray := []int{2, 4, 6, 8, 10}

	resultArray := make([]int, len(initArray))

	wg := sync.WaitGroup{}

	for i, num := range initArray {
		wg.Add(1)

		go func(i int, num int) {
			defer wg.Done()
			resultArray[i] = num * num
		}(i, num)
	}

	wg.Wait()

	fmt.Println("Initial array: ", initArray)

	fmt.Println("Result array: ", resultArray)
}
