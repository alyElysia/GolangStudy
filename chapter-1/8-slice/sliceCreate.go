package main

import (
	"fmt"
)

func sliceTest(slice1 []int) {
	for _, v := range slice1 {
		fmt.Println(v)
	}
}

func main() {

	// 1.slice的四种声明方式
	slice1 := []int{1, 1, 2, 5}

	var slice2 []int //与 slice11 := []int{} 等效

	var slice3 []int = make([]int, 3)

	slice4 := make([]int, 4) //最常用

	// ===================================================

	fmt.Printf("slice1's len = %d\n", len(slice1))
	fmt.Println("================")

	// 2.第二种声明方式仅声明slice2为一个切片，并没有为其分配存储空间，由于没有存储空间，因此无法为其赋值
	fmt.Printf("slice2'len = %d\n", len(slice2))
	// slice2[0] = 1 //会报错
	// 此时需要使用make来开辟空间
	slice2 = make([]int, 4) //第一个参数表示类型，第二个参数表示开辟的空间大小
	slice2[0] = 1
	fmt.Println("================")
	fmt.Printf("slice3's len = %d\n", len(slice3))

	fmt.Println("================")
	fmt.Printf("slice4's len = %d\n", len(slice4))

	// 3.判断一个slice是否为空（存储空间）
	if slice1 == nil {
		fmt.Println("slice1为空")
	} else {
		fmt.Println("slice1不为空")
	}
}
