package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	a = 10
	b = 20
	c = a + b

	// 只可以打印出格式化的字符串，可以输出字符串类型的变量，不可以输出整型变量和整型
	fmt.Printf("结果a = %d, b = %d , c = %d", a, b, c)
	fmt.Println("")
	// 可以打印出格式化的字符串，可以输出字符串类型的变量，不可以输出整型变量和整型
	fmt.Println("结果a = %d, b = %d , c = %d", a, b, c)
}
