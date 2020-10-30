package main

import "fmt"

/**
函数中可以定义多个函数

实现者不需要手动声明implement interface, 只要实现了函数, 就可以因为是一个实现者.
*/
type Phone interface {
	call(calledParty string)
}

// iPhone
type iPhone struct {
}

func (p iPhone) call(calledParty string) {
	fmt.Println("iPhone: Calling ", calledParty)
}

// Huawei
type Huawei struct {
}

func (p Huawei) call(calledParty string) {
	fmt.Println("Huawei: Calling ", calledParty)
}

func main() {
	// 实例化
	var p Phone
	// p.call("xx") // runtime error: invalid memory address or nil pointer dereference

	//new(iPhone).call("C99")
	//new(Huawei).call("XiaoYang")
	p = new(iPhone)
	p.call("C99")

	p = new(Huawei)
	p.call("XiaoYang")
}
