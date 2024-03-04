package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.WaitGroup 是一种用于等待一组 goroutine 完成执行的同步原语。它允许你协调和同步多个 goroutine 的执行
// WaitGroup 主要由三个方法组成：
// 			1.Add(delta int)：添加任务数。
// 			2.Wait()：阻塞等待所有任务的完成。
// 			3.Done()：完成任务。

// 函数默认为值传递，需要将sync.WaitGroup的指针传入
func Shopping(name string, wait *sync.WaitGroup) {
	fmt.Println(name, "正在购物~")
	time.Sleep(1 * time.Second)
	fmt.Println(name, "购物结束")
	// 在执行完毕时减去一个数量
	wait.Done()
}

func main() {

	t := time.Now()

	var wait sync.WaitGroup
	myWifes := []string{"洛琪希", "希露菲", "艾莉丝"}

	// 输入协程的数量
	wait.Add(len(myWifes))

	for _, v := range myWifes {
		go Shopping(v, &wait)
	}

	// 由于main结束后，其余协程都会结束，因此通过wait方法阻塞等待所有协程运行完毕后再执行main之后的部分
	wait.Wait()
	fmt.Println("花费时间：", time.Since(t))
}
