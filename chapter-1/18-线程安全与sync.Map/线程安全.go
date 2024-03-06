package main

import (
	"fmt"
	"sync"
)

// 通过为协程加锁来保障线程安全
var sum int
var wait sync.WaitGroup

// 定义一个锁
var lock sync.Mutex

func add() {
	// 上锁
	lock.Lock()
	for i := 0; i < 100000; i++ {
		sum++
	}
	wait.Done()
	// 解锁
	lock.Unlock()
}

func sub() {
	lock.Lock()
	for i := 0; i < 100000; i++ {
		sum--
	}
	wait.Done()
	lock.Unlock()
}

func main() {
	wait.Add(2)
	go add()
	go sub()
	wait.Wait()
	// 在不给协程加上锁的情况下sum的结果不会为0
	fmt.Println(sum)
}
