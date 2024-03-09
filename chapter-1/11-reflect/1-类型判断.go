package main

import (
	"fmt"
	"reflect"
)

func GetType(t any) {
	//1.类型判断
	reType := reflect.TypeOf(t) //TypeOf(v)获取v的具体类型

	//Kind()用于获取reType.Kind()的基础类型
	//具体类型相当于一个自定义类型、基础类型相当于自定义类型所属的类型
	switch reType.Kind() {
	case reflect.String:
		fmt.Println("这是一个String类型~")
	case reflect.Int:
		fmt.Println("这是一个Int类型~")
	case reflect.Struct:
		fmt.Println("这是一个Struct类型~")
	default:
		fmt.Println("未知类型...")
	}
}

func GetValue(v any) {
	//2.获取变量的值
	reValue := reflect.ValueOf(v)
	fmt.Println(reValue, reValue.Type())

	switch reValue.Kind() {
	case reflect.String:
		fmt.Println("这是一个String类型~", reValue.String())
	case reflect.Int:
		fmt.Println("这是一个Int类型~", reValue.Int())
	case reflect.Struct:
		fmt.Println("这是一个Struct类型~")
	default:
		fmt.Println("未知类型...")
	}
}

// 3.修改值
func SetValue(v any, v2 any) {
	rv1 := reflect.ValueOf(v)
	rv2 := reflect.ValueOf(v2)

	fmt.Println(rv1.Kind(), rv1.Elem(), rv1.Elem().Kind())

	//rv1.Elem()的含义是获取这个指针指向的值的reflect.value
	if rv1.Elem().Kind() != rv2.Kind() {
		fmt.Println("类型不匹配")
		return
	}

	switch rv1.Elem().Kind() {
	case reflect.String:
		rv1.Elem().SetString(rv2.String())
	case reflect.Int:
		rv1.Elem().SetInt(rv2.Int())
	}

}

func main() {
	//GetType(1)
	//GetType("12")
	//GetType(struct {
	//	name string
	//}{"lbw"})
	//
	//GetValue(1)
	//GetValue("12")
	//GetValue(struct {
	//	name string
	//}{"lbw"})

	name := "lbw"
	age := 18

	//修改值需要提供指针
	SetValue(&age, 19)
	SetValue(&name, "洛琪希")
	fmt.Println(name, age)

}
