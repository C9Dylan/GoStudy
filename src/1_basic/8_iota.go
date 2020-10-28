package main

import (
	"fmt"
)

/**
	iota: 可以理解为一个在编译期间就能确定的自增量。
 */
const (
	A = 1 << iota // iota = 0
	B = 1 << iota // iota = 1
	C			  // iota = 2
	D			  // iota = 3
)

func main() {
	fmt.Println(A, B, C, D)
}