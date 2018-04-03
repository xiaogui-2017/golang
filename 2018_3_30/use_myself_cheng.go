package main

import (
	"fmt"
)

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	// 0的阶乘为1
	return 1
}

func main() {
	var i int = 20
	// 强制类型转换
	fmt.Printf("%d的阶乘是%d\n", i, Factorial(uint64(i)))
}
