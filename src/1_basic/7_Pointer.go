package main

import "fmt"

func main() {

	/**
		1.运算符:

		&variable: 获取变量的指针
		*variablePointer: 通过变量指针值获取变量值
	 */

	// 2. 指针的类型如何声明
	pointerType()
	// 3. 通过指针取值, 改值案例
	pointerExample()
	// 4. 通过变量的指针来交换两个局参的值
	v1 := 10
	v2 := 20
	fmt.Print("v1: ", v1, "\n")
	fmt.Print("v2: ", v2, "\n")
	swap(&v1, &v2)
	fmt.Print("v1: ", v1, "\n")
	fmt.Print("v2: ", v2, "\n")
}

func pointerType()  {
	// 2. 指针type: var pointerName *variableType = &variable
	i := 10
	var i32 int32 = 10
	f := 9.9

	var iPointer *int = &i
	var i32Pointer *int32 = &i32
	var fPointer *float64 = &f

	fmt.Print(*iPointer, "\n")
	fmt.Print(*i32Pointer, "\n")
	fmt.Print(*fPointer, "\n")
}

func pointerExample() {
	// 声明变量
	variable := 10
	// 变量指针
	variablePointer := &variable
	// 通过指针获取的变量值
	variableValue := *variablePointer

	// 1. 打印指针16进制地址
	fmt.Printf("variable pointer address: %p\n", variablePointer)
	// 2. 打印通过指针获取的变量值
	fmt.Printf("variable value: %d\n", variableValue)
	// 3. 通过指针修改变量值后打印
	*variablePointer = 20
	fmt.Printf("variable value after modify: %d\n", variable)
}

func swap(p1, p2 *int)  {
	var t = *p1
	*p1 = *p2
	*p2 = t
}
