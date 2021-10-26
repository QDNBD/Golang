package main

import (
"fmt"
"reflect"
)

type Person2 struct {
}

func (p Person2)SelfIntroduction(name string, age int)  {
	fmt.Printf("Hello, my name is %s and i'm %d years old.", name, age)
}



func main() {
	p := &Person2{}

	//t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)
	name := reflect.ValueOf("wangbm")
	age := reflect.ValueOf(27)
	input := []reflect.Value{name, age}
	v.MethodByName("SelfIntroduction").Call(input)
}
