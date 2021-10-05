package main
import (
	"fmt"
	"time"
)
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, (i+1)*100)
	}
}
func say2(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Println(s, (i+1)*150)
	}
}
func main() {
	go say2("world")
	say("hello")
}

/*
问题来了，say2 只执行了 3 次，而不是设想的 5 次，为什么呢？

原来，在 goroutine 还没来得及跑完 5 次的时候，主函数已经退出了。

我们要想办法阻止主函数的结束，要等待 goroutine 执行完成之后，再退出主函数：
 */