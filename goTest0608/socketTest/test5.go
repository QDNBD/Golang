package main

import "fmt"

type reader interface {
	read() string
}

type writer interface {
	write() string
}

type rw interface {
	reader
	writer
}

type mouse struct{}

func (m mouse) read() string {
	return "mouse reading..."
}

func (m *mouse) write() string {
	return "mouse writing..."
}

func main() {
	var rw1 rw
	// 只要有一个指针实现，则此处必须是指针
	rw1 = &mouse{}
	// rw1 = new(mouse)
	fmt.Println(rw1.read())
	fmt.Println(rw1.write())
}