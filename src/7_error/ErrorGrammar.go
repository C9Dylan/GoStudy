package main

import "fmt"

/**
panic   -> throw
recover -> catch (异常只能被捕获一次, 再次补货则为nil)
defer   -> finally (可以存在多个defer, 但是需要提前在异常发生处之前声明.
					否则发生异常无法走defer, 且如果有多个, 按照声明顺序, 逆序执行)
*/
func main() {
	f1()
}

func f1() {
	fmt.Println("1")
	defer func() {
		error := recover()
		fmt.Println(error)
	}()
	defer handleError()
	f2()
	fmt.Println("4") // Don't exec
}

func f2() {
	fmt.Println("5")
	panic("Happen error")
	fmt.Println("6") // Don't exec
}

func handleError() {
	fmt.Println("2")
	error := recover()
	error2 := recover()
	fmt.Println(error)
	fmt.Println(error2)
	fmt.Println("3")
}
