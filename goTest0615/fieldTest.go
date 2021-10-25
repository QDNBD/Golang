package main

import (
	"fmt"
	"reflect"
)

type  Person struct {
	name string
	age  int
	gender string
}

func (p Person) SayBye() {
	fmt.Println("Bye")
}

func (p Person) SayHello() {
	fmt.Println("Hello")
}

func main()  {
	p :=  Person{"wanglong", 22, "male"}

	v := reflect.ValueOf(p)

	fmt.Println(v.NumField())
	fmt.Println(v.Field(0))
	fmt.Println(v.Field(1))
	fmt.Println(v.Field(2))

	fmt.Println("===========================")

	for i := 0; i < v.NumField(); i++ {
		fmt.Println(i+1, "    ", v.Field(i))
	}

	fmt.Println("method==========================")

	t := reflect.TypeOf(p)

	fmt.Println(t.NumMethod())
	fmt.Println(t.Method(0).Name)
	fmt.Println(t.Method(1).Name)

	fmt.Println("===================")

	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(i+1, "     " , t.Method(i).Name)
	}

	fmt.Println("===================")

	for i:=0; i<t.NumMethod(); i++{
		fmt.Printf("调用第 %d 个方法：%v ，调用结果：%v\n",
			i+1,
			t.Method(i).Name,
			v.Elem().Method(i).Call(nil))
	}

}