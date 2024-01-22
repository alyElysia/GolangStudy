package main

import "fmt"

func deferTestA() {
	fmt.Println("dA called...")
}

func deferTestB() {
	fmt.Println("dB called...")
}

func returnFunc() int {
	fmt.Println("return Func called...")
	return 0
}
func deferFunc() {
	fmt.Println("defer Func called...")
}

func deferAndReturn() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	// 1.被defer修饰的语句会在函数结束之后再执行
	defer fmt.Println("defer called...")
	fmt.Println("normal called...")

	// 2.被defer修饰的语句会放入栈中，因此当有多个语句被defer修饰后，执行的顺序与定义的顺序是相反的
	defer deferTestA()
	defer deferTestB()

	// 3.由于defer是在函数执行完毕时再调用的，因此当一个函数同时存在defer与return时，先执行return再执行defer
	deferAndReturn()
}
