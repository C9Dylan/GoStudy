package main

import "fmt"

/**
成员方法:

与函数不同, 函数可以直接调用.
成员方法需要通过: 特定类型变量.func 的形式调用.
*/
// 1. 声明一个类型
type Integer int

// 2. 针对于改类型, 绑定函数
func (intPtr Integer) print() {
	fmt.Println(&intPtr)
	fmt.Println(intPtr)
}

func (intPtr *Integer) toString() string { // 注意这里传递是指针, 因为Integer的元类型是int, 所以如果不传递指针, 无法改值
	fmt.Println(intPtr)
	return fmt.Sprint("Value: ", *intPtr)
}

/**
	🤔️ 如果元类型为引用类型, 如struct, 不使用 (s *theStruct), 而是(s theStruct), 在通过s去修改结构体中的字段值.

	这时会影响到调用者中的字段值吗? ---- No, 就像结构体作为方法形参时一样. 传递过来的都是副本.

    总结下来就是: 调用成员函数时, 会将调用者作为形参传递到成员函数中. 但是可以在成员方法上声明接收的是副本值, 还是调用者本身(指针).
*/

func main() {

	var i Integer = 10
	fmt.Println(&i)
	// 3. 通过变量.方法的方式调用
	i.print()
	fmt.Println(i.toString())
}
