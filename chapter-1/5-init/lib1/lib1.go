package lib1

import "fmt"

// 提供对外使用的API的名称首字母要大写
func Lib1Test() {
	fmt.Println("lib1 API ...")
}

func init() {
	fmt.Println("lib1 initing ...")
}
