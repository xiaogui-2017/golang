package main

import (
	"fmt"
)

func main() {
	// slice求和
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	// map的键值对中使用range
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range也可以枚举unicode字符串
	for i, c := range "hello world" {
		fmt.Println(i, )ci
	}
}
