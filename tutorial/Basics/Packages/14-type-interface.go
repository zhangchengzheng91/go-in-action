/*
类型推导

在声明一个变量而不指定其类型的时候(即使用不带类型的 := 语法或 var = 表达式)，变量的值由右值推导得出。
*/
package main

import "fmt"

func main() {
	v := "42"
	j := v
	f := 3.142
	g := 0.687 + 0.5i
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("j is of type %T\n", j)
	fmt.Printf("j is of type %T\n", f)
	fmt.Printf("j is of type %T\n", g)
}
