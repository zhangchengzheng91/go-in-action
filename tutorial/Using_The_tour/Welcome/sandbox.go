package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Welcome to playground\n")

	// 这里 gofmt 会包 warning，需要修改成下面的书写方式
	fmt.Printf("this itme is ", time.Now(), "\n")

	// 这里输出的时间格式不正确，后续有时间再调整优化
	fmt.Printf("this itme is %d", time.Now())
}
