package main

import "fmt"

func main() {
	/**
		声明多个变量的方式
	 */
	// 1
	var (
		n  = 1
		n2 = 1
	)
	fmt.Println(n, "\t", n2)

	// 2
	var n3, n4 = 2, 3
	fmt.Println(n3, "\t", n4)

	// 3
	var n5, n6 int32 = 4, 5
	fmt.Println(n5, "\t", n6)

	// 4
	n7, n8 := 7, 8
	fmt.Println(n7, "\t", n8)

	// Golang中, 可以利用多变量同时赋值的形式来校验两个参数值, 不使用临时变量T的情况下
	var n9 = 9
	var n10 = 10
	n9, n10 = n10, n9
	fmt.Println(n9, "\t", n10)
}