package main

import "fmt"

type Weekday int // 定义一个类型为Weekday, 且元类型为int

const (
	MONDAY Weekday = iota
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
)

func getWeekdayName(weekday Weekday) string {
	switch weekday {
		case MONDAY: return "Monday"
		case TUESDAY: return "Tuesday"
		case WEDNESDAY: return "Wednesday"
		case THURSDAY: return "Thursday"
		case FRIDAY: return "Friday"
		case SUNDAY: return "Sunday"
		default: return "None"
	}
}

func main() {

	/**
		枚举： golang并没有枚举类型数据, 但是可以通过iota进行模拟

		1. type 语法(类型别名): type 类型名称 实际类型(元类型, 即自定义类型的真实类型)

		2. iota: 第一个元素值为0, 下面元素的值默认+1, 因为默认使用套用第一个iota表达式.
				 也可以使用表达式来指定自增步长.
	 */

	fmt.Print(MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY, "\n")
	var temp Weekday = SUNDAY
	fmt.Println(temp)

	fmt.Println(getWeekdayName(SUNDAY))
	fmt.Println(getWeekdayName(8))
}
