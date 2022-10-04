/*
常量

常量只能使用 const 关键字声明，不能使用 var, :=

常量可以是字符，字符串，布尔值或数值
*/
package main

import "fmt"

const Pi = 3.14

var a = 9

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true

	fmt.Println("Go rules?", Truth)
	fmt.Println("this num is", a)
}
