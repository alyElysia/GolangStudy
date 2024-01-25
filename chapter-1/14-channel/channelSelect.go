package main

import "fmt"

// 通过实现斐波那契数列演示select监控多路channel的功能
func fibonacii(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			//如果c可以写入数据，那么执行以下代码
			x = y
			y += x
		case <-quit:
			//如果通道quit中有数据，那么执行以下代码
			fmt.Println("quit~")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		quit <- 0 //在执行循环后向quit写入数据
	}()

	fibonacii(c, quit)
}
