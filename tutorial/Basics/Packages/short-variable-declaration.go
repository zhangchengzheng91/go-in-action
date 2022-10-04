// 短变量声明
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// := 短变量声明语法，可以省略 var 和 数据类型
	// 函数外的每个语句都必须以关键字开始（var, func）,所以 := 不能在函数体外使用
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
