package bigNumbers

import (
	"fmt"
	"math/big"
)

func BigNumbersDemo() {
	fmt.Println("The operations of addition, subtraction, multiplication and division of two large numbers")

	numbers := NewBigNumbers(
		big.NewInt(1_000_000_000_000_000_000),
		big.NewInt(2_000_000_000_000_000_000),
	)

	fmt.Printf("Numbers: %d, %d\n", numbers.a, numbers.b)

	fmt.Println("Addition:   ", numbers.Add())
	fmt.Println("Subtraction:  ", numbers.Sub())
	fmt.Println("Multiplication:  ", numbers.Mul())

	div, err := numbers.DivRat()

	if err != nil {
		fmt.Println("Division: ", err)
	} else {
		fmt.Println("Division: ", div.FloatString(10))
	}
}
