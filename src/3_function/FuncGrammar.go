package main

import "fmt"

// 0. 声明一个函数为type:
type f1 func() // 无参
type f2 func(int) // 有参: 形参可以不用写名称
type f3 func(i int) // 有参: 形参可以不用写名称
type f4 func() int // 含有返回值
type f5 func() (int, string) // 含有返回值

type callback func(int)

func main() {
	// 1. 无返回值的函数: func fName (param type) {}
	println("CQQ")
	// 2. 单返回值的函数: func fName (param type) returnType {}
	fmt.Println(sum(10, 10))
	// 3. 多返回值的函数: r1, r2 := func fName (param type) returnType1, returnType2 {}
	arr, sum := multiply(1, 2, 3)
	for _, v := range arr {
		fmt.Print(v, "\t")
	}
	fmt.Println(sum)
	// 4. 函数作为参数: 参数中的函数该怎么写就怎么写
	handle(1, 2, func(i int) {
		fmt.Println(i)
	})
	// 5. 函数作为返回值:
	rf := returnFunc()
	i, s := rf(1)
	fmt.Println(i, s)
	fmt.Println(returnFunc()(2))
}

func println(p string) {
	fmt.Println(p)
}

func sum(x, y int) int {
	return x + y
}

func multiply(i, j, k int) ([]int, int) {
	return []int {i, j, k}, i * j * k
}

func handle(x, y int, cbf callback) {
	cbf(sum(x, y))
}

func returnFunc() func(i int) (j int, s string) {
	return func(i int) (j int, s string) {
		return i, fmt.Sprint("str:", i)
	}
}