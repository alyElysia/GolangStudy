package main

import "fmt"

func main() {
	c := make(chan int, 4)

	go func() {
		for i := 0; i < 4; i++ {
			c <- i
		}

		// 关闭channel
		close(c)
		// 注意：
		// 1.不能向被关闭的通道发送数据
		// 2.可以从被关闭的通道中读取数据
		// 3.向空通道（就是仅定义，未使用make声明）发送或读取数据都会发送阻塞
	}()

	// for {
	// 	// data, ok := <-c; ok的含义是，获取通道c中的数据以及状态，
	// 	// 如果ok为true表示通道未关闭，然后再将ok作为判断条件
	// 	if data, ok := <-c; ok { // 如果通道已经被关闭并且通道中的所有数据都被接收完，那么ok将会是false
	// 		fmt.Println(data)
	// 	} else {
	// 		break
	// 	}
	// }

	// channel与range
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Main Finished ...")
}
