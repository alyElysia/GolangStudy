package main

import "fmt"

func main() {
	// map的三种定义方式

	// 1.
	var map1 map[string]string //与slice相似，在使用map前必须调用make为其设置存储空间
	// 为map1开辟存储空间
	map1 = make(map[string]string) //可以不指定空间大小
	// fmt.Println(len(map1))
	// 为map1赋值
	map1["one"] = "java"
	map1["two"] = "c"
	map1["three"] = "c++"
	fmt.Println(map1) //map的输出顺序与定义顺序并不一致

	// 2.
	map2 := make(map[string]string, 3)
	map2["one"] = "java"
	map2["two"] = "c"
	map2["three"] = "c++"
	fmt.Println(map2)

	// 3.
	map3 := map[int]string{
		1: "java",
		2: "go",
		3: "c", //在最后一个元素后也要加上 ,
	}

	fmt.Println(map3)

	// tips:map只有len没有cap
}
