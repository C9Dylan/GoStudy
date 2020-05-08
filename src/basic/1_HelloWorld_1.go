// main包下的main函数才能执行, 否则无法执行!!!
package main

import "fmt"

func main() {
	fmt.Println("Hello World" + "Hello Go")
	/**
		两个命令:

		1: 直接运行 -> go run fileX.go

		2: 编译为shell -> go build fileX.go -> ./fileX.exe
	 */
}