package main

import (
"fmt"
"reflect"
)

type Person1 struct {
	name string
	age int
}

func (p Person1)SayBye() string {
	return "Bye"
}

func (p Person1)SayHello() string {
	return "Hello"
}

func (p Person1)SayBye1()  {
	fmt.Print("Bye1")
}

func (p Person1)SayHello1()  {
	fmt.Println("Hello1")
}

func main() {
	p := &Person1{"wangbm", 27}

	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)


	for i:=0;i<v.NumMethod();i++{
		fmt.Printf("调用第 %d 个方法：%v ，调用结果：%v\n",
			i+1,
			t.Method(i).Name,
			v.Elem().Method(i).Call(nil))
	}

	fmt.Println("==================")

	v.MethodByName("SayHello1").Call(nil)
	v.MethodByName("SayBye1").Call(nil)
}
