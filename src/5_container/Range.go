package main

import "fmt"

func main() {

	// 1. 数组/切片
	intArr := []int{1, 2, 3}
	for i, v := range intArr {
		fmt.Println(i, "\t", v)
	}
	println()

	// 2. String, s为字符在编码表中的数值
	str := "ABC"
	for i, s := range str {
		fmt.Println(i, "\t", s)
	}
	println()

	// 3. Map
	m := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}
	for k, v := range m {
		fmt.Println(k, "\t", v)
	}
}
