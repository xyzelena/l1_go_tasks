package main

import (
	"fmt"
	"l1/human"
	"l1/squaring"
)

func main() {
	Task1()
}

func Task1() {
	fmt.Println()
	fmt.Println("<-- Task 1 Human. Example with embedded struct -->")

	h := &human.Human{}
	h.SetName("Elena")
	h.SayHello()

	a := &human.Action{}
	a.SetName("Ivan")
	a.SetAge(45)

	a.Introduce()

	a.SayHello()

	fmt.Println()
	fmt.Println("<-- Task 2 Squaring. Example with goroutines -->")

	fmt.Println()
	fmt.Println("1. Variant with channels:")
	squaring.Squaring1()

	fmt.Println()
	fmt.Println("2. Variant with wait group:")
	squaring.Squaring2()
}
