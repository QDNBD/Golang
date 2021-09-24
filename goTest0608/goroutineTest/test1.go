package main

import "time"

func say(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		println(s)
	}
}

func main()  {
	go say("hello")
	say("world")
}
