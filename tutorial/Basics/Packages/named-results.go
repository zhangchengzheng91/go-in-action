package main

import "fmt"

// 直接返回变量
// 为了可读性，只能用在较短的函数当中
// 名称必须是有意义的
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))

}
