package main

import (
	"fmt"
)

func main() {
	A(1, 2, 3, 4, 5, 6)
}

/*
 有...类型自动变成slice(python中的列表),
 ...必须放在最后面， 否则语法错误
*/

// int和string只是传递值的拷贝，相当于浅拷贝
// slice传值， 地址的拷贝， 深拷贝， 对其本身进行操作
func A(a ...int) {
	fmt.Println(a)
}
