package concurrentEntry

import (
	"fmt"
	"sync"
)

var shoppingList = make(map[string]int)
var mutex sync.RWMutex

func AddToShoppingList(item string, count int, wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock()
	defer mutex.Unlock()

	shoppingList[item] = count

	fmt.Printf("Added %s: %d\n", item, count)
}

func RemoveFromShoppingList(item string, wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock()
	defer mutex.Unlock()
    
	if _, ok := shoppingList[item]; !ok {
		fmt.Println("Item not found: ", item)
		return
	}

	if shoppingList[item] > 1 {
		shoppingList[item]-=1
	} else {
		delete(shoppingList, item)
	}

	fmt.Println("Removed: ", item)
}

func ConcurrentEntryMutex() {
	wg := &sync.WaitGroup{}

	fmt.Println("Example with mutex")
	fmt.Println("Adding items to shopping list")

	wg.Add(5) 
	
	go AddToShoppingList("bread", 1, wg)
	go AddToShoppingList("milk", 2, wg)
	go AddToShoppingList("bread", 1, wg)
	go RemoveFromShoppingList("milk", wg)
	go RemoveFromShoppingList("bread", wg)

	wg.Wait() 
	
	mutex.RLock()
	fmt.Println("Final shopping list:", shoppingList)
	mutex.RUnlock()
}