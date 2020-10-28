package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/**
		1.变量声明:

		Format: var name type | var name (一般都是使用简写: var name, 编译器会自动根据右值类型进行参数类型推导)

		Example:
			var name string = "张三";

			var age int8 = 20;

		2.批量声明, 定义全局变量:

		Example:
			var (
				name type = val
				name2 type = val
			)
	 */

	/**
		3.短变量声明, 定义局部变量: name:=value(参数名必须是一个没有定义过的参数名称)
	 */
	age := 100
	fmt.Print(age)
	fmt.Print("\n")

	/**
		4. 基本数据类型(其中的数值代表bit)

		bool
		string
		int、int8、int16、int32、int64 // 有符号位的int. 范围: -128 ~ 127. int默认为32/64, 即int32或int64. 取决于CPU的位数
		uint、uint8、uint16、uint32、uint64、uintptr // 无符号位的int(unsigned int), 范围: 0 ~ 255. 默认为unit32/64, 即uint32或uint64. 取决于CPU的位数
		byte // uint8 的别名
		rune // int32 的别名 代表一个 Unicode 码
		float32、float64 // 默认根据CPU的位数分配变量的内存空间(32/64bit)
		complex64、complex128 // 复数

		当一个变量被声明之后，系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，
		指针为 nil 等。所有的内存在 Go 中都是经过初始化的。
	 */
	boolVal := true
	string := "string"

	fmt.Print("boolVal size: ", unsafe.Sizeof(boolVal), "\n")
	fmt.Print("string size: ",unsafe.Sizeof(string), "\n\n")

	var uintVal uint = 1
	var uint8Val uint8 = 1
	var uint16Val uint16 = 1
	var uint32Val uint32 = 1
	var uint64Val uint64 = 1
	fmt.Print("uintVal size: ",unsafe.Sizeof(uintVal), "\n")
	fmt.Print("uint8Val size: ",unsafe.Sizeof(uint8Val), "\n")
	fmt.Print("uint16Val size: ",unsafe.Sizeof(uint16Val), "\n")
	fmt.Print("uint32Val size: ",unsafe.Sizeof(uint32Val), "\n")
	fmt.Print("uint64Val size: ",unsafe.Sizeof(uint64Val), "\n\n")

	var intVal int = 1
	var int8Val int8 = 1
	var int16Val int16 = 1
	var int32Val int32 = 1
	var int64Val int64 = 1
	fmt.Print("intVal size: ",unsafe.Sizeof(intVal), "\n")
	fmt.Print("int8Val size: ",unsafe.Sizeof(int8Val), "\n")
	fmt.Print("int16Val size: ",unsafe.Sizeof(int16Val), "\n")
	fmt.Print("int32Val size: ",unsafe.Sizeof(int32Val), "\n")
	fmt.Print("int64Val size: ",unsafe.Sizeof(int64Val), "\n\n")

	var floatVal = 9.9
	var float32Val float32 = 9.9
	var float64Val float64 = 9.9
	fmt.Print("floatVal size: ",unsafe.Sizeof(floatVal), "\n")
	fmt.Print("float32Val size: ",unsafe.Sizeof(float32Val), "\n")
	fmt.Print("float64Val size: ",unsafe.Sizeof(float64Val), "\n")
	const f1 = 9.9e2 // 9.1 * 10 ^ 2
	const f2 = 9.9e-2 // 9.1 / 10 ^ 2
	const f3 = 9.9E2 // 9.1 * 10 ^ 2
	const f4 = 9.9E-2 // 9.1 / 10 ^ 2
	const f5 = 1. // 小数后可省略不写或者干脆不写小数位
	const f6 float64 = 2
	const f7 = .99 // 99 / 10 ^ 2
	fmt.Print(f1, "\t", f2, "\t", f3, "\t", f4, "\t", f5, "\t", f6, "\t", f7, "\n")


	var byteVal byte = 1
	var runeVal rune = 1
	fmt.Print("byteVal size: ",unsafe.Sizeof(byteVal), "\n")
	fmt.Print("runeVal size: ",unsafe.Sizeof(runeVal), "\n\n")

	/**
		5. 空白表示符号: _

		用于舍弃部分不需要的函数返回值, 或者变量值.
	 */
	_, emptyFlagVariable := 1, 10
	fmt.Println(emptyFlagVariable)

	r1, r2 := emptyFlagTestFunc(
		1,
		2,
		"3",
	)
	fmt.Println(r1, r2)

	_, r2 = emptyFlagTestFunc(
		0,
		0,
		"4",
	)
	fmt.Println(r2)
}

func emptyFlagTestFunc(v1, v2 int, v3 string) (int, string) {
	return v2, v3
}
