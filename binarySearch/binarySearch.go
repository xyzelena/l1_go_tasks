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
	return bsRec(arr, searchItem, 0, len(arr)-1)
}

func bsRec(arr []int, searchItem int, left, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2

	if arr[mid] == searchItem {
		return mid
	}

	if arr[mid] < searchItem {
		return bsRec(arr, searchItem, mid+1, right)
	}

	return bsRec(arr, searchItem, left, mid-1)
}
