package main

import (
	"fmt"
	"time"
)

/*
我们引入一个信道，默认的，信道的存消息和取消息都是阻塞的，
在 goroutine 中执行完成后给信道一个值 0，则主函数会一直等待信道中的值，一旦信道有值，主函数才会结束。
 */

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, (i+1)*100)
	}
}
func say2(s string, ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Println(s, (i+1)*150)
	}
	ch <- 0
	close(ch)
}

func main() {
	ch := make(chan int)
	go say2("world", ch)
	say("hello")
	fmt.Println(<-ch)
}

