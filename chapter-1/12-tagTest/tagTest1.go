package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	// 在定义字段后加上标签来做出更详细的说明
	// 标签由一个或多个键值对组成
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", tagInfo, " doc: ", tagDoc)

	}
}

func main() {
	var re resume
	findTag(&re)
}
