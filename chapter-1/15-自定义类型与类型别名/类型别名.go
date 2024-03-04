package main

import "fmt"

// 类型别名与自定义类型的不同之处
// 1.不能绑定方法
// 2.与原类型的类型一致，只是额外设置了一个别名
// 3.设置别名类型的变量可直接与原类型一致的变量进行比较，而自定义类型需要进行转换

type mi = int
type dint int

func main() {
	var i1 mi = 11
	var i2 dint = 11
	var i3 int = 11

	fmt.Printf("i1的类型是：%T\n", i1)
	fmt.Printf("i2的类型是：%T\n", i2)

	if i1 == i3 {
		fmt.Println("类型别名与原类型比较")
	}
	if i2 == dint(i3) {
		fmt.Println("自定义类型与原类型比较")
	}
}
