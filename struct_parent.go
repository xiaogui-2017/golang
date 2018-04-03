package main

import (
	"fmt"
)

type A struct {
	B
}

type B struct {
	C
	Name string
}

type C struct {
	Name string
}

func main() {
	a := A{B: B{Name: "B"}}
	fmt.Println(a.B.C.Name)
}
