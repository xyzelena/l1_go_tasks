package deletingSliceElement

import (
	"errors"
	"fmt"
)

func DeletingSliceElement(slice []int, index int) ([]int, error) {
	fmt.Println("Initial slice:", slice)
	fmt.Println("Index to delete:", index)

	switch {
	case index < 0 || index >= len(slice):
		return nil, errors.New("index out of range")
	case len(slice) == 0:
		return nil, errors.New("slice is empty")
	default:
		copy(slice[index:], slice[index+1:])
		slice[len(slice)-1] = 0
		slice = slice[:len(slice)-1]
		return slice, nil
	}
}
