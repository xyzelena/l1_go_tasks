package wordReversal

import (
	"fmt"
	"strings"
)

func ReverseWords(s string) string {
	fmt.Println("Initial string: ", s)

	words := strings.Fields(s)

	if len(words) == 0 || len(words) == 1 {
		fmt.Println("Reversed string: ", s)
		return s
	} else if len(words) == 2 {
		words[0], words[1] = words[1], words[0]
	} else {
		for i := 0; i < len(words)/2; i++ {
			words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
		}
	}

	reversedString := strings.Join(words, " ")

	fmt.Println("Reversed string: ", reversedString)

	return reversedString
}
