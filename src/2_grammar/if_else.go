package main

import "fmt"

func main() {
	// if else
	checkAge(20)
	// elseif
	checkScore(85)
}

func checkAge(age int) {
	if age >= 18 {
		fmt.Println("成年")
	} else {
		fmt.Println("未成年")
	}
}

func checkScore(score int)  {
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
}
