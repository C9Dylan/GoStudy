package main

import (
	"fmt"
	"reflect"
)

func main() {

	// 1. 声明: map[keyType] valueType { key:value, ...}
	countries := map[int]string{
		1: "China",
		2: "England",
		3: "America",
		4: "Italy",
		5: "Kiwi",
		6: "Turkey",
	}

	printMap(countries)

	// 1) 增
	countries[7] = "France"
	printMap(countries)
	// 2) 查
	fmt.Println(countries[7])
	// 查询一个不存在的key, 返回valueType类型的默认值
	v := countries[8]
	t := reflect.TypeOf(v)
	fmt.Println(v == "")
	fmt.Println(t)
	// 3) 删除
	delete(countries, 7)
	printMap(countries)

	// 4) 改
	countries[1] = "China No.1"
	printMap(countries)
}

func printMap(m map[int]string) {
	for k, v := range m {
		fmt.Println(k, "\t", v)
	}
	println()
}
