package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	a := person{
		Name: "zhang_zhiyuan",
		Age:  19,
	}
	fmt.Println(a)
	// 调用修改结构体的函数， 结构体作参数传入
	A(a)
	fmt.Println(a)
}

func A(per person) {
	// 结构体是值传递，不能改变原来的值的内容
	per.Age = 100
	fmt.Println("func A ", per)
}
