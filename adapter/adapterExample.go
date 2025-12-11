package adapter

import "fmt"

func AdapterExample() {
	fmt.Println("Adapter Example")
	old := &OldInterface{}
	adapter := &Adapter{oldInterface: old}
	adapter.DoNewSomething("Hello, World!")
}
