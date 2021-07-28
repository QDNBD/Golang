package main

import "fmt"

func main() {
	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 appped() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)

	// Step 3: 显示两行数据
	fmt.Println("Row 1")
	fmt.Println(values[0])
	fmt.Println("Row 2")
	fmt.Println(values[1])

	// Step 4: 访问第一个元素
	fmt.Println("第一个元素为：")
	fmt.Println(values[0][0])
}


/*
初始化
a := [3][4]int{
 {0, 1, 2, 3} ,  // 第一行索引为 0
{4, 5, 6, 7} ,   //  第二行索引为 1
{8, 9, 10, 11}}   //第三行索引为 2


func main() {
    // 创建二维数组
    sites := [2][2]string{}

    // 向二维数组添加元素
    sites[0][0] = "Google"
    sites[0][1] = "Runoob"
    sites[1][0] = "Taobao"
    sites[1][1] = "Weibo"

    // 显示结果
    fmt.Println(sites)
}


func main() {
   //数组 - 5 行 2 列
	var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
	var i, j int

	//输出数组元素
	for  i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
		}
	}
}



func main() {
    // 创建空的二维数组
    animals := [][]string{}

    // 创建三一维数组，各数组长度不同
    row1 := []string{"fish", "shark", "eel"}
    row2 := []string{"bird"}
    row3 := []string{"lizard", "salamander"}

    // 使用 append() 函数将一维数组添加到二维数组中
    animals = append(animals, row1)
    animals = append(animals, row2)
    animals = append(animals, row3)

    // 循环输出
    for i := range animals {
        fmt.Printf("Row: %v\n", i)
        fmt.Println(animals[i])
    }
}


func main() {
    // 二维数组
    var value = [3][2]int{{1, 2}, {3, 4}, {5, 6}}
    // 遍历二维数组的其他方法，使用 range
    // 其实，这里的 i, j 表示行游标和列游标
    // v2 就是具体的每一个元素
    // v  就是每一行的所有元素
    for i, v := range value {
        for j, v2 := range v {
            fmt.Printf("value[%v][%v]=%v \t ", i, j, v2)
        }
        fmt.Print(v)
        fmt.Println()
    }
}


//这样是没问题的
func main() {
   var arr = [3][4]int{
      {0, 1, 2, 3} ,      {4, 5, 6, 7} ,      {8, 9, 10, 11}}
   fmt.Println(arr[1][1])
}

//把外层括号单独拿到下一层就会报错
func main() {
   var arr = [3][4]int{
      {0, 1, 2, 3} ,      {4, 5, 6, 7} ,      {8, 9, 10, 11}
   }
   fmt.Println(arr[1][1])
}

//把外层括号单独拿到下一层但是在结尾加上逗号就可以了
func main() {
   var arr = [3][4]int{
      {0, 1, 2, 3} ,      {4, 5, 6, 7} ,      {8, 9, 10, 11},   }
   fmt.Println(arr[1][1])
}
 */