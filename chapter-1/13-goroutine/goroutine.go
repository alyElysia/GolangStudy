package main

import (
	"fmt"
	"runtime"
	"time"
)

// 子协程
func newTask() {
	i := 0
	for {
		fmt.Println(" i = ", i)
		i++
	}
}

// 主协程
func main() {

	// go中的协程通过关键字go来标识

	// func() {
	// }()	这是一个匿名函数
	go func() {
		defer fmt.Println("A defer ...")
		func() {
			defer fmt.Println("B defer ...")

			runtime.Goexit() //退出当前协程
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	go newTask() //设置协程的另一种方式

	// 死循环，防止主协程结束导致从协程结束
	for {
		time.Sleep(1 * time.Second)
	}
}
