package main

import (
	"fmt"
)

// 1.接口的使用：接口不能绑定方法、保存的是值类型、存储 值+原始类型
// 实现接口的所有方法被称为实现对应接口
type SingInterface interface {
	Sing()
	GetName() string
}

type Chicken struct {
	Name string
}

type Cat struct {
	Name string
}

func (cc Chicken) Sing() {
	fmt.Println(cc.GetName(), "，在唱K")
}
func (cc Chicken) GetName() string {
	return cc.Name
}

func (ct Cat) Sing() {
	fmt.Println(ct.GetName(), "，在唱K")
}
func (ct Cat) GetName() string {
	return ct.Name
}

func sing(si SingInterface) {
	si.Sing()
	// fmt.Println(si.GetName())
}

// 2.类型的断言
func Check(si SingInterface) string {

	// 第一种用法：将对象判断为指定类型，返回该对象以及一个布尔值
	// t, ok := si.(Chicken)
	// fmt.Println(t, "是一只坤吗：", ok)

	// 第二种用法：指定为该对象的所有可能的类型
	switch si.(type) {
	case Chicken:
		return "这是一只坤坤"
	case Cat:
		return "这是一只咪咪"
	}

	return ""
}

// 3.空接口的使用
// 空接口可接受任何类型的值
// any是interface{}的别名
func PrintI(a any) {
	fmt.Println(a)
}

func main() {
	cc := Chicken{Name: "坤坤"}
	ct := Cat{Name: "咪咪"}
	sing(cc)
	sing(ct)
	fmt.Println(Check(cc))
	fmt.Println(Check(ct))
	PrintI(cc)
	PrintI(1)
}
