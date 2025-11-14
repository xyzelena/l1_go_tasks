package main

import (
	"fmt"
	"l1/concurrentEntry"
	"l1/exchangeValues"
	"l1/goroutineStops"
	"l1/human"
	"l1/intersectionSets"
	"l1/pipeline"
	"l1/quickSort"
	"l1/setStrings"
	"l1/settingBit"
	"l1/squaring"
	"l1/tempGrouping"
	"l1/timeoutChannel"
	"l1/variableType"
	"l1/workers"
)

func main() {
	// Task1()   //Human
	// Task2()   //Squaring
	// Task3_4() //Workers with correct completion by pressing Ctrl+C (SIGINT)
	// Task5()   //Timeout per channel
	// Task6()   //Goroutine Stops
	// Task7()   //Concurrent Entry
	// Task8()   //Setting Bit
	// Task9()   //Pipeline
	// Task10()  //Temperature Grouping
	// Task11()  //Intersection of Sets
	// Task12() //Set of unique strings
	// Task13() //Exchange of values ​​without a third variable
	// Task14() //Detect Variable Type
	Task15() //Quick Sort
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

func Task6() {
	fmt.Println()
	fmt.Println("<-- Task 6 Goroutine Stops -->")

	goroutineStops.GoroutineStops()
}

func Task7() {
	fmt.Println()
	fmt.Println("<-- Task 7 Concurrent entry in the map -->")

	concurrentEntry.ConcurrentEntryMutex()

	fmt.Println()

	concurrentEntry.ConcurrentEntryMap()
}

func Task8() {
	fmt.Println()
	fmt.Println("<-- Task 8 Setting Bit Operations -->")

	var x int64 = 5 // 0101₂
	fmt.Printf("Before setting bit: %d (%04b)\n", x, x)

	y, old := settingBit.SetBit(x, 1, 0)
	fmt.Printf("After setting 1-st bit to 0: %d (%04b), old bit = %d\n", y, y, old)

	z, old := settingBit.SetBit(x, 1, 1)
	fmt.Printf("After setting 1-st bit to 1: %d (%04b), old bit = %d\n", z, z, old)
}

func Task9() {
	fmt.Println()
	fmt.Println("<-- Task 9 Pipeline -->")

	pipeline.Pipeline()
}

func Task10() {
	fmt.Println()
	fmt.Println("<-- Task 10 Temperature Grouping -->")

	tempGrouping.TempGrouping()
}

func Task11() {
	fmt.Println()
	fmt.Println("<-- Task 11 Intersection of Sets -->")

	intersectionSets.IntersectionSets()
}

func Task12() {
	fmt.Println()
	fmt.Println("<-- Task 12 Set of unique strings -->")

	setStrings.SetStrings()
}

func Task13() {
	fmt.Println()
	fmt.Println("<-- Task 13 Exchange of values without a third variable -->")

	exchangeValues.ExchangeValuesWithAdditionAndSubtraction()
	exchangeValues.ExchangeValuesWithXOR()
	exchangeValues.ExchangeValues()
}

func Task14() {
	fmt.Println()
	fmt.Println("<-- Task 14 Detect Variable Type -->")

	variableType.DetectVariableType(1)
	variableType.DetectVariableType("Hello")
	variableType.DetectVariableType(true)
	variableType.DetectVariableType(make(chan int))
	variableType.DetectVariableType(nil)
	variableType.DetectVariableType([]string{"a", "b"})
	variableType.DetectVariableType(map[string]int{"a": 1, "b": 2})
}

func Task15() {
	fmt.Println()
	fmt.Println("<-- Task 15 Quick Sort -->")

	arr := []int{3, 6, 8, 10, 1, 2, 1}
	fmt.Println("Initial array:", arr)

	sortedArr := quickSort.QuickSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
