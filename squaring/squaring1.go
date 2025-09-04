// squaring with channels

package squaring

import "fmt"

type result struct {
	index int
	value int
}

func Squaring1() {
	initArray := []int{2, 4, 6, 8, 10}

	resultArray := make([]int, len(initArray))

	ch := make(chan result, len(initArray))

	for i, num := range initArray {
		go func(i int, num int) {
			ch <- result{index: i, value: num * num}
		}(i, num)
	}

	for range initArray {
		res := <-ch
		resultArray[res.index] = res.value
	}

	fmt.Println("Initial array: ", initArray)

	fmt.Println("Result array: ", resultArray)
}
