package main

import (
	"fmt"
)

func main() {
	var a [2]int
	var b [1]int
	b = a
	fmt.Println(b)
}
