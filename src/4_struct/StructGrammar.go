package main

import "fmt"

// 0. 声明结构体
type Person struct {
	// 成员变量
	name   string
	age    int
	salary float64
	family []Person
}

func main() {
	// 1. 创建
	p := createPerson()

	// 2. getter
	fmt.Println(p.name, "\t", p.age, "\t", p.salary)

	// 3. setter
	p.name = "C99-Upd"
	fmt.Println(p.name)
	setName(p, "C99-Upd2") // 不生效, 因为当结构体当参数传递给方法时, 就像传递基本类型数据一样, 传递当是一个副本. 所以需要传递指针.
	fmt.Println(p.name)
	setNameByPointer(&p, "C99-Upd3")
	fmt.Println(p.name)
}

func createPerson() Person {
	// 使用变量名创建
	return Person{
		name:   "C99",
		age:    20,
		salary: 20_000.0,
		// 也可以不使用变量名创建, 但是不能部分使用变量名, 部分不使用
		family: []Person{
			{"XiaoYang", 18, 200_000.0, []Person{ /* ... */ }},
		},
	}
}

func setName(p Person, name string) {
	p.name = name
	fmt.Println("In setName func: ", p.name)
}

func setNameByPointer(p *Person, name string) {
	p.name = name
}
