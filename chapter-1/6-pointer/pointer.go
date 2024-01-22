package main

import "fmt"

func swap(ta *int, tb *int) {
	var temp int
	temp = *ta
	*ta = *tb
	*tb = temp
}

func main() {
	a, b := 10, 20
	fmt.Println("a,b交换前的值分别为：", a, b)
	swap(&a, &b)
	fmt.Println("a,b交换后的值分别为：", a, b)
}
