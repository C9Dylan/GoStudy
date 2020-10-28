package main

func main() {

	// 闭包函数: 能够读取其他函数内部变量的函数.

	getId := getSequence() // 闭包函数: getId, 可以修改getSequence内部的计数变量

	for i:= 0; i < 5; i++ {
		print(getId(), "\t")
	}
}

func getSequence() func() int {
	i := 0
	return func () int {
		i++
		return i
	}
}
