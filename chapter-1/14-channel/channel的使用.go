package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var moneyChan = make(chan int)

func Shopping(name string, cost int, wait *sync.WaitGroup) {
	fmt.Println(name, "正在购物~")
	time.Sleep(1 * time.Second)
	fmt.Println(name, "购物结束,花费：", cost)
	moneyChan <- cost
	wait.Done()
}

func main() {

	t := time.Now()

	var wait sync.WaitGroup
	myWifes := []string{"洛琪希", "希露菲", "艾莉丝"}

	wait.Add(len(myWifes))

	for _, v := range myWifes {
		go Shopping(v, rand.Intn(10)+1, &wait)
	}

	// 这个协程用于等待上方所有的协程的完成
	go func() {
		defer close(moneyChan) //当上方所有的协程执行完毕之后关闭通道，如果不关闭通道该协程会一直阻塞进而发生死锁
		wait.Wait()

	}()

	var costs int
	// 这个for循环的作用是循环取出moneyChan中所有的值
	for cost := range moneyChan {
		costs += cost
	}
	fmt.Println("花费时间：", time.Since(t), "花费金额：", costs)
}
