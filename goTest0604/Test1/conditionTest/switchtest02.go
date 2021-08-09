package main

import (
	"fmt"
)

func findType(i interface{}) {
	switch x := i.(type) {
	case int:
		fmt.Println(x, "is int")
	case string:
		fmt.Println(x, "is string")
	case nil:
		fmt.Println(x, "is nil")
	case float64:
		fmt.Println(x, "is float64")
	default:
		fmt.Println(x, "not type matched")
	}
}

func main() {
	findType(10)          //int
	findType(2.487439734) //float64

	var val float32 = 3.45
	findType(val)

	findType(4.55)
	var k interface{} //nil
	findType(k)
	findType("hello")

}
