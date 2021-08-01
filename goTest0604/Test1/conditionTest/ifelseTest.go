package main

import "fmt"

func main(){
	// var count,c int   //定义变量不使用也会报错
	var count int
	var flag bool
	count=1
	//while(count<100) {    //go没有while
	for count < 100 {
		count++
		flag = true;
		//注意tmp变量  :=
		for tmp:=2;tmp<count;tmp++ {
			if count%tmp==0{
				flag = false
			}
		}

		// 每一个 if else 都需要加入括号 同时 else 位置不能在新一行
		if flag == true {
			fmt.Println(count,"素数")
		}else{
			continue
		}
	}
}

/**
package main

import"fmt"

func main(){
    var a int
    var b int
    fmt.Printf("请输入密码：   \n")
    fmt.Scan(&a)
    if a == 5211314 {
    fmt.Printf("请再次输入密码：")
    fmt.Scan(&b)
        if b == 5211314 {
            fmt.Printf("密码正确，门锁已打开")
        }else{
            fmt.Printf("非法入侵，已自动报警")
        }
    }else{
        fmt.Printf("非法入侵，已自动报警")
    }
}
 */