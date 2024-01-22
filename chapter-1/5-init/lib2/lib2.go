package lib2

import "fmt"

// 提供对外使用的API的名称首字母要大写
func Lib2Test() {
	fmt.Println("lib2 API ...")
}

func init() {
	fmt.Println("lib2 initing ...")
}
