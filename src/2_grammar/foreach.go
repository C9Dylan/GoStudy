package main

import "fmt"

func main() {
	// 1. 基本循环
	// 1) for condition {}
	asWhile(10)
	// 2) for init; condition; post {}
	foreach(5)
	// 3) for {}
	endlessLoop(3)
	// 4) for k,v := range container {}
	foreachContainer()

	// Goto keyword
	gotoKeyword(5)
}

func asWhile(n int) {
	c := 0
	for n > 0 {
		c += n
		n--
	}
	fmt.Println(c)
}

func foreach(n int) {
	for i := 0; i <= n; i++ {
		fmt.Print(i, "\t")
	}
	fmt.Println()
}

func endlessLoop(n int) {
	// for ;; {}
f:
	for {
		fmt.Print(n, "\t")
		n--
		if n < 0 {
			break f
		}
		// continue f
	}
	fmt.Println()
}

func foreachContainer() {

	arr := []string{"A", "B", "C"}

	for i, v := range arr {
		fmt.Println(i, "\t", v)
	}
}

func gotoKeyword(n int) {

	c := 0

label:
	fmt.Print(n, "\t")

	if n > 0 {
		c += n
		n--
		goto label
	}

	fmt.Println(c)
}
