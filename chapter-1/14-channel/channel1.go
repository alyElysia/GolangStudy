package main

import (
	"fmt"
	"time"
)

func main() {
	// 有缓存的通道

	c := make(chan int, 3) //定义一个容量为3的通道
	go func() {
		defer fmt.Println("goroutine over ...")

		fmt.Println("len(c) = ", len(c), "cap(c) = ", cap(c))

		// 在写入的数据量大于channel时会发生阻塞，直到数据被读出
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("写入通道的值为：", i, "len(c) = ", len(c), "cap(c) = ", cap(c))
		}
	}()

	time.Sleep(2 * time.Second) //防止主协程阻塞

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("从通道获取的值为：", num)
	}

	// 子协程此时处于阻塞状态，在主协程睡眠的过程中会重新执行
	time.Sleep(2 * time.Second) //防止主协程结束使得子协程被迫结束
	fmt.Println("main routine over ...")
}
