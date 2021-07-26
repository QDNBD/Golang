package main
/*
var arrays[10] int

var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

balances := [5]float32{100.0, 324.0, 323.9,  23.09, 432.43}

//如果数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度：
var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
或
balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

如果设置了数组的长度，我们还可以通过指定下标来初始化元素：
//  将索引为 1 和 3 的元素初始化
balance := [5]float32{1:2.0,3:7.0}


 */


import "fmt"

func main() {
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i,j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j] )
	}
}

/*
func main() {
   var i,j,k int
   // 声明数组的同时快速初始化数组
   balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

   // 输出数组元素
	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i] )
	}

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	//输出每个数组元素的值
	for j = 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j] )
	}

	//  将索引为 1 和 3 的元素初始化
	balance3 := [5]float32{1:2.0, 3:7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k] )
	}
}
*/
