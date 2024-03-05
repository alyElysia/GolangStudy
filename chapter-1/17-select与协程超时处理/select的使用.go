package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var moneyChan = make(chan int)
var nameChan = make(chan string)
var doneChan = make(chan struct{}) //作为一个信号通道

func Shopping(name string, cost int, wait *sync.WaitGroup) {
	fmt.Println(name, "正在购物~")
	time.Sleep(1 * time.Second)
	fmt.Println(name, "购物结束,花费：", cost)
	moneyChan <- cost
	nameChan <- name
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

	go func() {
		defer close(moneyChan)
		defer close(nameChan)
		defer close(doneChan) // 当所有协程执行完毕后，关闭该协程作为信号
		wait.Wait()

	}()

	var nameLists []string
	var costLists []int

	// 定义一个事件函数
	events := func() {
		// 循环读取通道中的数据
		for {
			select {
			case cost := <-moneyChan:
				costLists = append(costLists, cost)
			case name := <-nameChan:
				nameLists = append(nameLists, name)
			case <-doneChan: //当试图从一个关闭的通道中读取数据时，会返回对应类型的零值，而不会阻塞
				// 结束该函数
				return
			}
		}
	}
	events()
	fmt.Println("花费时间：", time.Since(t), "nameLists：", nameLists, "costLists：", costLists)
}
