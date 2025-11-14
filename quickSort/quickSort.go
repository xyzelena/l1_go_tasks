package quickSort

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]

	var left, right, equal []int

	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			equal = append(equal, v)
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, equal...), right...)
}
