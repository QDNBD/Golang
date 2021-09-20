package main

import (
	"fmt"
	"time"
)

func main() {
	pipline := make(chan int)

	go func() {
		fmt.Println("准备发送数据: 100")
		pipline <- 100
	}()

	go func() {
		num := <-pipline
		fmt.Printf("接收到的数据是: %d", num)
	}()
	// 主函数sleep，使得上面两个goroutine有机会执行
	time.Sleep(time.Second)
}



//定义只读信道

var pipline1 = make(chan int)
type Receiver1 = <-chan int // 关键代码：定义别名类型
var receiver Receiver = pipline1


//定义只写信道

var pipline2 = make(chan int)
type Sender1 = chan<- int  // 关键代码：定义别名类型
var sender1 Sender = pipline2

//定义只写信道类型
type Sender = chan<- int

//定义只读信道类型
type Receiver = <-chan int

func main1() {
	var pipline = make(chan int)

	go func() {
		var sender Sender = pipline
		fmt.Println("准备发送数据: 100")
		sender <- 100
	}()

	go func() {
		var receiver Receiver = pipline
		num := <-receiver
		fmt.Printf("接收到的数据是: %d", num)
	}()
	// 主函数sleep，使得上面两个goroutine有机会执行
	time.Sleep(time.Second)
}


//遍历信道，可以使用 for 搭配 range关键字，在range时，要确保信道是处于关闭状态，否则循环会阻塞。


func fibonacci(mychan chan int) {
	n := cap(mychan)
	x, y := 1, 1
	for i := 0; i < n; i++ {
		mychan <- x
		x, y = y, x+y
	}
	// 记得 close 信道
	// 不然主函数中遍历完并不会结束，而是会阻塞。
	close(mychan)
}

func main3() {
	pipline := make(chan int, 10)

	go fibonacci(pipline)

	for k := range pipline {
		fmt.Println(k)
	}
}