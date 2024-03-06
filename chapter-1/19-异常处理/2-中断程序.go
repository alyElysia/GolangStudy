package main

import (
	"os"
)

// 中断程序基本都是在项目初始化发生问题时使用
func init() {
	_, err := os.ReadFile("12354")
	if err != nil {
		// 通过下方两种方法将程序中断
		panic(err)
		// log.Fatalln(err)

	}
}

func main() {

}
