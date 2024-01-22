package main

import (
	"fmt"
)

// 数组的数据类型与长度相对应
func arrTest(arr [4]int) {
	// 2.for循环遍历数组的方式
	// 第一种方式
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 在调用数组作为参数时属于 值拷贝 ，无法修改原数组中的元素的值
	arr[0] = 110
}

// 动态数组的数据类型与长度不绑定
func sliceTest(slice1 []int) {
	// for循环遍历数组的第二种方式
	// 使用range关键字，其中第一个参数是index，表示对应数组元素的下标，第二个参数是value，表示对应的值
	// 如果不需要使用其中的一个参数，可以设置其为匿名，即 _ 。
	for _, v := range slice1 {
		fmt.Printf("value = %d\n", v)
	}

	// 在调用切片作为参数时属于引用传递，
	// 即调用的是切片的指针，因此可以对原切片的值进行修改
	slice1[0] = 111
}

func main() {

	// 1.数组的定义
	arr1 := [4]int{1, 23, 5, 4}
	arrTest(arr1)

	// 3.切片的定义：：slice := []int{} 只定义不赋值的格式，数组需要加上长度
	// 切片又称为动态数组
	slice1 := []int{1, 5, 8, 7}
	sliceTest(slice1)
	fmt.Println("=====输出修改后的slice1=======")
	for _, v := range slice1 {
		fmt.Println("value = ", v)
	}
}
