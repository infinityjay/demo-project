package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{"Alice", 30}

	val := reflect.ValueOf(person)
	typ := val.Type()

	// 遍历结构体的所有字段
	for i := 0; i < typ.NumField(); i++ {
		// 获取每个字段的结构体字段类型
		field := typ.Field(i)

		// 获取字段名和值
		fmt.Println(field.Name, val.Field(i).Interface())
	}
}
