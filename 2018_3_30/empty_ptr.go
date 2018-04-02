package main

import (
	"fmt"
)

func main() {
	var prt *int
	fmt.Printf("ptr的值是%x\n", prt)
	fmt.Printf("ptr的值是%d\n", prt)
}

//go指针还可以进行空指针判断
