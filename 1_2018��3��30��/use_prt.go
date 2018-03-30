package main

import (
	"fmt"
)

func main() {
	var a int = 20

	// * 号用于指定变量是作为一个指针
	var ip *int
	ip = &a
	fmt.Printf("a变量的地址是:%x\n", &a)
	fmt.Printf("ip变量储存的指针地址是:%x\n", ip)
	fmt.Printf("*ip变量的值：%d", *ip)

	// *相关的就是值
	// &相关的就是地址
}
