package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s3 := s1[:]
	s4 := s1[:len(s1)]
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}
