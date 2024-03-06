package main

import (
	"errors"
	"fmt"
)

// 将异常信息抛到向上层函数
func div(a, b int) (res int, err error) {
	if b == 0 {
		err = errors.New("除数不能为0")
		return
	}
	res = a / b
	return
}

func server() (res int, err error) {
	res, err = div(2, 0)
	if err != nil {
		//向上抛
		return
	}
	// 其他的逻辑...
	res++
	return
}

func main() {
	res, err := server()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
