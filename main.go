package main

import (
	"fmt"
	"l1/human"
	"l1/squaring"
	"l1/timeoutChannel"
	"l1/workers"
)

func main() {
	Task1()   //Human
	Task2()   //Squaring
	Task3_4() //Workers with correct completion by pressing Ctrl+C (SIGINT)
	Task5()   //Timeout per channel
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
}

func Task2() {
	fmt.Println()
	fmt.Println("<-- Task 2 Squaring. Example with goroutines -->")

	fmt.Println()
	fmt.Println("1. Variant with channels:")
	squaring.Squaring1()

	fmt.Println()
	fmt.Println("2. Variant with wait group:")
	squaring.Squaring2()
}

func Task3_4() {
	fmt.Println()
	fmt.Println("<-- Task 3 and 4 Workers. Example with N workers reading from channel -->")
	fmt.Println("<-- with completion by pressing Ctrl+C -->")

	countWorkers := 3
	fmt.Printf("Start %d workers...\n", countWorkers)

	workers.DoWorkers(countWorkers)
}

func Task5() {
	fmt.Println()
	fmt.Println("<-- Task 5 Timeout per channel -->")

	timeoutChannel.TimeoutChannel()
}
