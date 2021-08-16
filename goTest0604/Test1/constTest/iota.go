package main

import "fmt"

//iota，特殊常量，可以认为是一个可以被编译器修改的常量。
//ota 只是在同一个 const 常量组内递增，每当有新的 const 关键字时，iota 计数会重新开始。

func main() {
	const (
		a = iota   //0
		b          //1
		c          //2
		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
	)
	const a1 = iota

	fmt.Println(a,b,c,d,e,f,g,h,i,a1)
}