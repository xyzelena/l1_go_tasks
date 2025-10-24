package intersectionSets

import (
	"fmt"
	"slices"
)

func IntersectionSets() {
	fmt.Println("Intersection of sets")

	set1 := []int{1, 2, 3, 4}
	set2 := []int{3, 4, 5}

	fmt.Println("Set 1:", set1)
	fmt.Println("Set 2:", set2)

	intrSet := make([]int, 0)

	for _, value := range set1 {
		if slices.Contains(set2, value) {
			intrSet = append(intrSet, value)
		}
	}

	fmt.Println("Intersection of sets:", intrSet)
}
