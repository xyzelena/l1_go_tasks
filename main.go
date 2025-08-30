package main

import (
	"fmt"
	"l1/human"
)

func main() {
	Task1()
}

func Task1() {
	fmt.Println("<-- Task 1 Human. Example with embedded struct -->")

	h := &human.Human{}
	h.SetName("Elena")
	h.SayHello()

	a := &human.Action{}
	a.SetName("Ivan")
	a.SetAge(45)

	a.Introduce()

	a.SayHello()
}