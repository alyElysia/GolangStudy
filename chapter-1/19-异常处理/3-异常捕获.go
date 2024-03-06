package main

import (
	"fmt"
	"runtime/debug"
)

func read() {
	defer func() {
		// 调用recover来捕获panic，防止程序中断
		err := recover()
		fmt.Println(err)
		// 打印当前的调用栈来找到错误的位置
		fmt.Println(string(debug.Stack()))
	}()

	lists := []int{1, 3, 4}
	fmt.Println(lists[3])
}

func main() {
	read()
	fmt.Println("正常逻辑....")
}
