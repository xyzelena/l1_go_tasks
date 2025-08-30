package human

import "fmt"

type Human struct {
	name string
}

func (h *Human) SetName(name string) {
	h.name = name
}

func (h Human) GetName() string {
	return h.name
}

func (h Human) SayHello() {
	fmt.Println("Hello, my name is ", h.GetName())
}
