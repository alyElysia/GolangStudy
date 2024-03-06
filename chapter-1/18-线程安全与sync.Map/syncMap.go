package main

import (
	"fmt"
	"sync"
)

// 与线程安全类似，当在对一个map同时进行读写操作时也会发生安全问题
func main() {
	// var map1 = map[int]string

	var map2 = sync.Map{} // 底层实现了上锁去锁的操作，适用于并发场景

	// 线程不安全的情况：两个协程在同时对map进行读写操作
	// go func() {
	// 	for i := 0; i < 1000000; i++ {
	// 		map1[i] = "洛琪希"
	// 	}
	// }()
	// go func() {
	// 	for i := 0; i < 1000000; i++ {
	// 		fmt.Println(map1[i])
	// 	}
	// }()

	go func() {
		for {
			map2.Store(1, "洛琪希")
		}
	}()

	go func() {
		for {
			v, _ := map2.Load(1)
			fmt.Println(v)
		}
	}()

	// 在协程未执行完毕前都在此阻塞
	select {}
}
