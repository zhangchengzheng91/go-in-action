package main

import "fmt"

// 函数能够接收 0 个或 多个 参数
func add(x int, y int) int {
	return x + y
}

// 如果两个参数的数据类型相同，在最后给定参数的数据类型
func add2(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(43, 13))
	fmt.Println(add2(43, 13))
}
