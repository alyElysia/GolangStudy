package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("goruntine is over... ")
		c <- 123 //为管道赋值
	}()

	num, ok := <-c //接收管道的值，并判断管道是否关闭，为true表示为关闭
	fmt.Println("在从协程中通过channel传递的值为 ", num)
	if ok {
		fmt.Println("管道未关闭")
	}
	// c <- 12
	fmt.Println("main goroutine is over ...")
}
