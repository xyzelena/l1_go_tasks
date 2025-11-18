package binarySearch

func BinarySearchWithLoop(arr []int, searchItem int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == searchItem {
			return mid
		} else if arr[mid] < searchItem {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func BinarySearchWithRecursion(arr []int, searchItem int) int {
	if len(arr) == 0 {
		return -1
	}

	mid := len(arr) / 2

	if arr[mid] == searchItem {
		return mid
	}

	if arr[mid] < searchItem {
		return BinarySearchWithRecursion(arr[mid+1:], searchItem)
	}

	return BinarySearchWithRecursion(arr[:mid], searchItem)
}