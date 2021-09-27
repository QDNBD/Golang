package main

import "fmt"

func main()  {
	//sum := 0
	//for i := 1; i < 10; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)

	//sum := 1
	//for ; sum <= 10; {
	//	sum += sum
	//}
	//fmt.Println(sum)
	//
	//for sum <= 20 {
	//	sum += sum
	//}
	//fmt.Println(sum)

	//sum := 0
	//for {
	//	sum++ // 无限循环下去
	//}
	//fmt.Println(sum) // 无法输出
	//
	//strings := []string{"google", "runoob"}
	//for i, s := range strings {
	//	fmt.Println(i, s)
	//}
	//
	//
	//numbers := [6]int{1, 2, 3, 5}
	//for i,x:= range numbers {
	//	//fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	//	fmt.Println(i,x)
	//}


	//九九乘法表
	//for m := 1; m < 10; m++ {
	//	/*    fmt.Printf("第%d次：\n",m) */
	//	for n := 1; n <= m; n++ {
	//		fmt.Printf("%dx%d=%d ",n,m,m*n)
	//	}
	//	fmt.Println("")
	//}

	//下实例使用循环嵌套来输出 2 到 100 间的素数：
	/* 定义局部变量 */
	var i, j int

	for i=2; i < 100; i++ {
		for j=2; j <= (i/j); j++ {
			if(i%j==0) {
				break; // 如果发现因子，则不是素数
			}
		}
		if(j > (i/j)) {
			fmt.Printf("%d  是素数\n", i);
		}
	}
}
