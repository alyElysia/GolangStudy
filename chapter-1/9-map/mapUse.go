package main

import "fmt"

func printMap(map1 map[string]string) {
	for k, v := range map1 {
		fmt.Printf("key = %s, value = %s\n", k, v)
	}
}

// 2.map的调用
func changeMap(Tmap1 map[string]string) {
	// map作为参数时也属于引用传递，可以对被调用的map进行修改
	for k, v := range Tmap1 {
		Tmap1[k] = k + "'s capital is " + v
	}
}

func main() {

	map1 := make(map[string]string)

	// 1.map的增删改查
	// 增加
	map1["China"] = "BeiJing"
	map1["Japan"] = "Tokyo"
	map1["USA"] = "NewYork"

	// 删除
	// delete(map1, "USA") //指定要删除的map和对应键

	// 修改
	// map1["Japan"] = "DABAN"

	//查询
	printMap(map1)

	fmt.Println("===============")
	changeMap(map1)
	printMap(map1)
}
