package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title       string  // 1. 可以被转为Json
	price       float64 // 2. 无法被转为Json, 因为字段不是public
	Author      string  `json:"author"`      // 3. 将被转换的JsonKey设定为指定值. 避免JsonKey一直都为大写的结构体中的字段名称.
	contentType string  `json:"ContentType"` // 4. 验证`json:"xxx"`只是适用于toJsonString时
}

func main() {
	book := Book{"数据结构", 100.0, "严蔚敏", "数据结构"}
	boonJsonByte, err := json.Marshal(book)
	fmt.Println(string(boonJsonByte), "\n", err)
}
