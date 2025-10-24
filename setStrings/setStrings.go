package setStrings

import (
	"fmt"
)

func SetStrings() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println("Strings:", strings)

	setUniqueStrings := make(map[string]struct{})

	for _, s := range strings {
		setUniqueStrings[s] = struct{}{}
	}

	fmt.Println("Set of unique strings:")

	for s := range setUniqueStrings {
		fmt.Printf("%s ", s)
	}
}
