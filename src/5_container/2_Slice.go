package main

import "fmt"

func main() {

	// 1. 创建一个切片
	// 1) 数组也可以理解为一个切片, 数组默认capacity == length
	intArr := []int{
		1,
		2,
		3,
	}
	// 2) 通过make创建切片
	makeIntArr := make([]int, 3, 5)

	// 3) 通过切片符: `:`, 注意: 仅仅是将数组的引用赋值给了copyIntArr. 通过copyIntArr是可以修改intArr中的元素的
	copyIntArr := intArr[:]
	copyIntArr[0] = -1

	printlnSlice(intArr)
	printlnSlice(makeIntArr)
	printlnSlice(copyIntArr)

	// 2. 创建空切片
	var emptySlice []int
	printlnSlice(emptySlice)
	fmt.Println(emptySlice, "\t", "emptySlice equal nil ? ", emptySlice == nil)

	// 3. 查看切片的长度与容量: len & cap

	/**
	4. 切片截取:

	slice[startIndex:] -> 指定索引到最后一个元素 | 指定长度后到所有数据
	slice[:endIndex] -> 第一个元素开始, 截止到指定索引-1处 | 倒数y个长度元素之前到所有元素
	slice[startIndex:endIndex] -> 指定起始索引开始, 截止到指定结尾索引-1处
	*/
	sliceArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printlnSlice(sliceArr[5:])
	printlnSlice(sliceArr[:5])
	printlnSlice(sliceArr[5:7])

	// 测试切片与源数组是否为同一个: 是
	s := sliceArr[0:2]
	s[0] = -1
	printlnSlice(s)
	printlnSlice(sliceArr)

	// 5. append
	mi := make([]int, 1)
	printlnSlice(mi)
	mi = append(mi, 1)
	printlnSlice(mi)
	mi = append(mi, 2)
	printlnSlice(mi)
	// 5.1 测试append的返回值: 会返回一个新的切片, 并不会将添加元素追加到mi中
	printlnSlice(append(mi, 3))
	printlnSlice(append(mi, 4)) // 没有3
	// 6. copy
	ci := []int{1, 2, 3}
	mci := make([]int, 2)
	copy(mci, ci)
	printlnSlice(mci)
}

func printlnSlice(intArr []int) {
	fmt.Println(intArr, "\t", "Length: ", len(intArr), "\t", "Capacity: ", cap(intArr))
}
