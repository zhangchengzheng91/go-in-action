package main

import "fmt"

// 多变量声明，如果数据类型相同，可以将数据类型放到最后
// var 声明变量，可以在函数内，也可以在函数外
var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
