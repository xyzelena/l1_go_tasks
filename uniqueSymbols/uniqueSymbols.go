package uniqueSymbols

import (
	"errors"
	"fmt"
	"strings"
)

func CheckUniqueSymbols(str string) (bool, error) {
	fmt.Println("String: ", str)

	if len(str) == 0 {
		fmt.Println(len(str))
		return false, errors.New("Пустая строка!")
	}

	lowerStr := strings.ToLower(str)

	strMap := make(map[rune]int)

	for _, char := range lowerStr {
		strMap[char] += 1

		if strMap[char] > 1 {
			fmt.Println("Символы не уникальны!")
			return false, nil
		}
	}

	fmt.Println("Символы уникальны!")
	return true, nil
}
