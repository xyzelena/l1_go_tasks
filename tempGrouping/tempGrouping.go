package tempGrouping

import "fmt"

func TempGrouping() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	tempGroup := make(map[int][]float64)

	for _, t := range temp {
		tempGroup[int(t/10)*10] = append(tempGroup[int(t/10)*10], t)
	}

	fmt.Println("Initial temperature:")
	fmt.Println(temp)

	fmt.Println("Temperature grouping:")
	fmt.Println(tempGroup)
}
