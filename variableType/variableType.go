package variableType

import (
	"fmt"
	"reflect"
)

func DetectVariableType(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Println("Значение:", val, "Тип: int")
	case string:
		fmt.Println("Значение:", val, "Тип: string")
	case bool:
		fmt.Println("Значение:", val, "Тип: bool")
	case chan int, chan string, chan bool, chan interface{}:
		fmt.Println("Значение:", val, "Тип: chan (канал)")
	default:
		t := reflect.TypeOf(val)
		if t == nil {
			fmt.Println("Значение:", val, "Тип: nil")
		} else {
			fmt.Printf("Значение: %v. Тип: %s\n", val, t)
		}
	}
}
