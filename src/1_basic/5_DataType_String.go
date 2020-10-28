package main

import "fmt"

func main() {

	/**
		String专栏

		编码: Golang中字符串使用Unicode的实现UTF8编码(ASCII的字符一个字节, ASCII编制之外的根据字符内容使用2 ~ 4个字符存储, 即 一个字符占用1 ~ 4个byte).

	 	转义字符::

		%v 打印结构体
		%+V 打印带有字段的结构体
		%T 打印对象类型
		%t 打印布尔值
		%d 打印整型数，十进制输出，如果d前面有数字，表示控制输出宽度，默认使用空白填充，%05d 会在不满5位时填充0
		%b 打印整型数，二进制输出
		%c 打印整型数，字符输出（如果有）
		%o 打印整型数，八进制输出，如果x前面带有#表示带有零的八进制
		%x 打印整型数，十六进制输出，如果x前面带有#表示带有0x的十六进制
		%f 打印浮点数，正常输出，如果f前面有x.y 表示控制宽度和输出小数点位数，要对其，再加-，例如 %-10.5f
		%e,%E 打印浮点数，科学记数法输出
		%s 打印字符串
		%q 打印字符串，带有引号输出
		%x 打印字符串，使用base-16输出其编码
		%p 打印指针
		%U 打印Unicode字符
		%#U 打印带字符的Unicode
	 */

	/**
		1. 字符串拼接
	 */
	// 1) + 号连接
	name := "Kim" + "-" + "张三"
	fmt.Printf(name + "\n")
	// b) 大文本: ``(数字键1的左边, esc的下面), 且将包含的内存全部当做字符串处理, 即类似\t | \n这样的转移字符无效.
	description := `
		嘤嘤嘤\r嘤嘤嘤
		嘤嘤嘤
	`
	fmt.Print(description)
	fmt.Println()
	/**
		2. 字符串的迭代
	 */
	str := "Welcome to Northeast! 哒哒哒"
	for x := 0; x < len(str); x++ {
		fmt.Printf("%c %d\n", str[x], str[x]) // 对于在ASCII表中的字符可以使用str[index]的形式获取, 但是ASCII编码表之外的字符获取后是识别不了的.
	}
	fmt.Println()
	for index, s := range str {
		fmt.Printf("%d %c\n", index, s)
	}
}
