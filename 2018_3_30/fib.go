/*
先分析，分情况
*/
package main

import (
	"fmt"
)

// 圆括号后的int是返回值的类型
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t\t", fib(i))
	}
}
