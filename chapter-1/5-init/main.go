package main

import (

	// 导包测试
	// 在未使用模块化时需要将环境变量中的GO111MODULE的值设为off才可正常使用
	// "chapter-1/5-init/lib1"
	// 可以为导入的包取别名，其中：包名为 _ ，表示匿名，即使未调用也会自行初始化，
	// 包名为 . ，在调用时对应包中的方法时不需要加上包名，直接写方法名即可，不推荐使用
	// "chapter-1/5-init/lib1"
	"fmt"
)

func funcTest(a int, b int) (c int, d int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c, d = 11, 22
	return
}

func main() {
	// fmt.Println("Hello,world")

	// fmt.Println("-----------------")
	// fmt.Println(funcTest(11, 22))
	// fmt.Println("-------------")
	// re1, re2 := funcTest(11, 22)
	// fmt.Println(re1, re2)
	var s1 string
	fmt.Println("s1 的默认值为：", s1)
	// lib1.Lib1Test()

}
