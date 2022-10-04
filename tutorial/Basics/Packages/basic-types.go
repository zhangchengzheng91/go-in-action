// 基本类型
/*
bool // false

string // ""

int int8[byte] int16 int32[rune] int64 // 0
uint uint8 uint15 uint32 unit64 uintptr // 0

float32 float64 // 0

complex64 complex128 // 0

*/
package main

import (
	"fmt"
	"math/cmplx"
)

// 同导入语句一样，变量声明也可以分组成一个语法块
var (
	Tobe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", Tobe, Tobe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
