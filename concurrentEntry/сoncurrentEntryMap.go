package concurrentEntry

import (
	"fmt"
	"sync"
)

func addToLessonsList(lesson, teacher string, wg *sync.WaitGroup, lessonsList *sync.Map) {
	defer wg.Done()

	lessonsList.Store(lesson, teacher)

	fmt.Println("Added:", lesson, "by", teacher)
}

func removeFromLessonsList(lesson string, wg *sync.WaitGroup, lessonsList *sync.Map) {
	defer wg.Done()

	teacher, ok := lessonsList.Load(lesson)
	if !ok {
		fmt.Println("Lesson not found:", lesson)
		return
	}

	lessonsList.Delete(lesson)

	fmt.Println("Removed:", lesson, "by", teacher)
}

func ConcurrentEntryMap() {
	fmt.Println("Example with concurrent-map")

	lessonsList := &sync.Map{}

	var wg sync.WaitGroup

	wg.Add(5)
	go addToLessonsList("Math", "John Doe", &wg, lessonsList)
	go addToLessonsList("History", "Jane Smith", &wg, lessonsList)
	go addToLessonsList("Physics", "Jim Beam", &wg, lessonsList)
	go removeFromLessonsList("Math", &wg, lessonsList)
	go removeFromLessonsList("History", &wg, lessonsList)
	wg.Wait()

	fmt.Println("Final lessons list:")

	lessonsList.Range(func(key, value any) bool {
		fmt.Printf("%s taught by %s\n", key.(string), value.(string))
		return true
	})
}
