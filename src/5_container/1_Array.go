package main

import "fmt"

func main() {

	/**
	1. 声明:

		var arrayName [size]elementType

		var arrayName  = [] elementType {element....}

		var arrayName  = [size] elementType {element....}

	2. 多纬数组: [][][]...
	*/
	var emptyArray [3]int
	var arr1 = [3]int{1, 2, 3}
	var arr2 = []string{"A", "B", "C"} // 根据元素数量推断数组长度
	fmt.Println(emptyArray)
	fmt.Println(arr1)
	fmt.Println(arr2)
}
