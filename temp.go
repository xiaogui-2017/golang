package main

import (
	"fmt"
	"reflect"
)

// 类声明
type User struct {
	Id   int
	Name string
	Age  int
}

// 类方法
func (u User) Hello() {
	fmt.Println("hello world")
}

func main() {
	// 类的实例化对象
	u := User{1, "ok", 12}
}
