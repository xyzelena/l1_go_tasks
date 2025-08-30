package human

import "fmt"

type Action struct {
	Human // Встроенная структура (embedded struct) - аналог наследования
	age int
}

func (a *Action) SetAge(age int) {
	a.age = age
}

func (a Action) GetAge() int {
	return a.age
}

func (a Action) Introduce() {
	fmt.Println("Hi! I'm ", 
	a.GetName(),
	" and my age is: ", 
	a.GetAge())
}
