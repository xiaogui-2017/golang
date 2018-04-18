package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("hello")
}
func main() {
	u := User{1, "zhangzhiyuan", 25}
	Info(u)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o) // User
	fmt.Println("type:", t.Name())

	v := reflect.ValueOf(o) // {1 zhangzhiyuan 25}
	//fmt.Print(t.NumField())  // 3
	//fmt.Print(v.NumField())  // 3

	for i := 0; i < v.NumField(); i++ {
		// 遍历结构体
		fmt.Printf("%16v%16s%16v\n",v.Field(i),t.Field(i).Name, t.Field(i).Type)
	}
}
