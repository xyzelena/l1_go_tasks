package adapter

import "fmt"

type OldInterface struct{}

func (old *OldInterface) DoOldSomething(text string) {
	fmt.Println("From Old Interface:", text)
}
