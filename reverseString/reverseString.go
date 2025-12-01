package reverseString

import "fmt"

func ReverseString(s string) string {
	r := []rune(s)

	fmt.Println("Initial string: ", s)

	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-i-1] = r[len(r)-i-1], r[i]
	}

	fmt.Println("Reversed string: ", string(r))

	return string(r)
}
