package main

import (
	"fmt"
)

// 定义了一个接口phone,包含call()方法
type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, i can call you")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am IPhone, i can call you")
}

func main() {
	// 定义了一个Phone类型的phone变量， 接口就是类型， 实例化后就和方法绑定
	var phone Phone
	var iphone Phone
	// 相当于类型实例化后才能调用call方法
	phone = new(NokiaPhone) // new一个结构
	phone.call()

	iphone = new(IPhone)
	iphone.call()
}
