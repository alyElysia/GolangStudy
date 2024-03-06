package main

import (
	"fmt"
	"time"
)

var tChan = make(chan struct{})

func send() {
	fmt.Println("正在发送文件~")
	time.Sleep(2 * time.Second)
	fmt.Println("文件发送完毕")
	close(tChan)
}

func main() {
	go send()

	select {
	case <-tChan:
		fmt.Println("文件发送完毕~")
	case <-time.After(1 * time.Second): //设置最大限制时长
		fmt.Println("发送超时")
		return
	}
}
