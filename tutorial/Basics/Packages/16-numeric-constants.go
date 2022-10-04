package main

import "fmt"

const (
	// 将 1 左移 100 位来创建一个非常大的数
	// 即这个数是二进制1后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，，即 Small = 1 << 1 或 Small = 2
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
