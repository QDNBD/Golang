package main

import (
	"fmt"
	"strings"
)

func main()  {
	str := "这 里 是\nww w  .qin iu.\ncom\n"
	fmt.Println("这里是最初的字符串")
	fmt.Println(str)
	fmt.Println("去除空格")
	str = strings.Replace(str, " ", "", -1)
	fmt.Println(str)
	fmt.Println("去除换行")
	str = strings.Replace(str, "\n", "", -1)
	fmt.Println(str)

}
