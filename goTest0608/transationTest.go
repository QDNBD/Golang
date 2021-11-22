package main

import "fmt"
//不支持隐式转换
func main() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum)/float32(count)
	fmt.Printf("mean 的值为: %f\n",mean)
	var b int32
	b = int32(mean)
	fmt.Println(b)
}
