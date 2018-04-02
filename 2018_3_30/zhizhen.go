// Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
package main

import (
	"fmt"
)

func main() {
	var a int = 10
	fmt.Printf("变量的地址是:%x \n", &a)
}

// * 号用于指定变量是作为一个指针
// var ip *int 指向整型
// var fp *float32 指向浮点型
