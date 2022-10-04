/*
类型转换

表达式 T(v) 将值 v 转换为类型 T。

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = uint(f)

	fmt.Println(x, y, z)

	i := 42
	ft := float64(i)
	u := uint(ft)
	fmt.Println(i, ft, u)
}
