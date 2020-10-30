package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
注: 此文件内容需要先学习: 4_struct
*/

// 定义性别枚举
type Gender int

const (
	MALE Gender = iota + 1
	FEMALE
)

func (gPtr *Gender) toString() string {
	switch *gPtr {
	case MALE:
		return "Men"
	case FEMALE:
		return "Women"
	default:
		return "Unknown"
	}
}

type Person struct {
	Name   string
	Age    int
	gender Gender
}

func (p *Person) generatePersonalInformation() string {
	var builder strings.Builder // 1. 实例化结构体
	builder.WriteString("Name: ")
	builder.WriteString(p.Name) // 2. 使用绑定在结构体上的函数, 追加的元素会被存储在结构体中
	builder.WriteString("\t")
	builder.WriteString("Age: ")
	builder.WriteString(strconv.Itoa(p.Age))
	builder.WriteString("\t")
	builder.WriteString("Gender: ")
	builder.WriteString(p.gender.toString())
	return builder.String()
}

func main() {
	p := Person{"C99", 20, MALE}
	information := p.generatePersonalInformation()
	fmt.Println(information)
}
