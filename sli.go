package main

import (
	"fmt"
)

func main() {
	// var sli []int  声明一种空的slice
	// fmt.Println(sli)

	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	s1 := a[5:len(a)]
	s2 := a[5:]
	s3 := a[:5]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	// 如果10不够，每次增加一倍，容量已知， 提升效率
	s4 := make([]int, 3, 10)
	fmt.Println(s4) //跟我们使用数组是一样的
	fmt.Println(len(s4), cap(s4))

	b := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}
	sb := b[2:5]
	fmt.Println(string(sb))
}
