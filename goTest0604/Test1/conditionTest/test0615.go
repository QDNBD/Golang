// package main

// import "fmt"

// func main() {
// 	any := make([]interface{}, 4)

// 	any[0] = 11
// 	any[1] = "String hello"
// 	any[2] = []int{1, 2, 3, 4, 5}
// 	for _, value := range any {
// 		fmt.Println(value)
// 	}

// }

//package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	var age interface{} = 25

// 	fmt.Printf("原始接口变量的类型为 %T，值为 %v \n", age, age)

// 	t := reflect.TypeOf(age)
// 	v := reflect.ValueOf(age)

// 	// 从接口变量到反射对象
// 	fmt.Printf("从接口变量到反射对象：Type对象的类型为 %T \n", t)
// 	fmt.Printf("从接口变量到反射对象：Value对象的类型为 %T \n", v)

// 	// 从反射对象到接口变量
// 	i := v.Interface()
// 	fmt.Printf("从反射对象到接口变量：新对象的类型为 %T 值为 %v \n", i, i)

// }

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	var name string = "Go编程时光"

// 	v := reflect.ValueOf(&name)
// 	fmt.Println("可写性为:", v.CanSet())
// }

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	var name string = "Go编程时光"
// 	v1 := reflect.ValueOf(&name)
// 	fmt.Println("v1 可写性为:", v1.CanSet())

// 	v2 := v1.Elem()
// 	fmt.Println("v2 可写性为:", v2.CanSet())
// }

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	var name string = "Go编程时光"
// 	fmt.Println("真实世界里 name 的原始值为：", name)

// 	v1 := reflect.ValueOf(&name)
// 	v2 := v1.Elem()

// 	v2.SetString("Python编程时光")
// 	fmt.Println("通过反射对象进行更新后，真实世界里 name 变为：", name)
// }

package main

import (
	"fmt"
	"reflect"
)

func appendToSlice(arrPtr interface{}) {
	valuePtr := reflect.ValueOf(arrPtr)
	value := valuePtr.Elem()

	value.Set(reflect.Append(value, reflect.ValueOf(3)))
	value.Set(reflect.Append(value, reflect.ValueOf(4)))

	fmt.Println(value)
	fmt.Println(value.Len())
}

func main() {
	arr := []int{1, 2}

	appendToSlice(&arr)

	fmt.Println(arr)
}
