package exchangeValues

import "fmt"

func ExchangeValuesWithAdditionAndSubtraction() {
	fmt.Println("Exchange values with addition and subtraction")

	a, b := 10, 20

	fmt.Println("a =", a, "b =", b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("a =", a, "b =", b)
}

func ExchangeValuesWithXOR() {
	fmt.Println("Exchange values with XOR")

	a, b := 5, 7

	fmt.Println("a =", a, "b =", b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println("a =", a, "b =", b)
}

func ExchangeValues() {
	fmt.Println("Exchange values with multiple assignment")

	a, b := 3, 8

	fmt.Println("a =", a, "b =", b)

	a, b = b, a

	fmt.Println("a =", a, "b =", b)
}
