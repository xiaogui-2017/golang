package main

import (
	"fmt"
)

type TZ int

func (tz *TZ) increace(num int) {
	*tz += TZ(num)
}

func main() {
	var a TZ
	a.increace(100)
	fmt.Println(a)
}
