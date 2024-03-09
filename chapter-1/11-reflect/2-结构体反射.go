package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Student struct {
	Name  string `json:"name"`
	Age   int
	IsMan bool
}

func (s Student) Call(name string) {
	fmt.Println(name, "调用了这个方法")
}

// PersonInfo 1.获取结构体的值
func PersonInfo(std any) {
	t := reflect.TypeOf(std)
	v := reflect.ValueOf(std)

	//NumField returns a struct type's field count.
	for i := 0; i < t.NumField(); i++ {
		// Field returns a struct type's i'th field.
		tf := t.Field(i)

		//判断字段的tag是否为空
		jsonTag := tf.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = tf.Name //将字段名赋值给tag
		}

		fmt.Println(jsonTag, v.Field(i))
	}
}

// SetStruct 2.修改结构体字段的值
func SetStruct(obj any) {
	v := reflect.ValueOf(obj).Elem()
	t := reflect.TypeOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		Stag := t.Field(i).Tag.Get("json")
		if Stag != "" {
			value.SetString(strings.ToUpper(value.String()))
		}
	}
}

// UseMethod 3.调用结构体的方法
func UseMethod(obj any) {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	//NumMethod()方法的作用是返回这个结构体所包含的方法数量
	for i := 0; i < v.NumMethod(); i++ {

		m := t.Method(i) //获取方法的名称，类型等元信息
		if m.Name == "call" {
			continue
		}

		method := v.Method(i) //获取实际可调用的方法

		//调用该方法
		method.Call([]reflect.Value{
			//传递参数
			reflect.ValueOf("洛琪希"),
		})
	}
}

func main() {
	std := Student{
		Name:  "lqx",
		Age:   18,
		IsMan: false,
	}
	SetStruct(&std)
	PersonInfo(std)
	UseMethod(std)

}
