//在定义常量组时，如果不提供初始值，则表示将使用上行的表达式。


package main

import "fmt"

const (
	a = 1
	b
	c
	d
)

func main() {
	fmt.Println(a)
	// b、c、d没有初始化，使用上一行(即a)的值
	fmt.Println(b)   // 输出1
	fmt.Println(c)   // 输出1
	fmt.Println(d)   // 输出1
}
