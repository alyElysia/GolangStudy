package main

import (
	"encoding/json"
	"fmt"
)

// 定义一个fx变量存储多个类型
type fx interface {
	int | int8 | int16 | int32 | uint | uint8 | uint16 | uint32
}

// 1.泛型函数
// 在函数名后加上[]，其中输入一个字符占位，再写入该字符可代表的类型
func FxFunc[T fx](a, b T) T {
	return a + b
}

// 2.泛型结构体
type Users[T any] struct {
	Name     string `json:"name"`
	Msg      string `json:"msg"`
	UserInfo T      `json:"userInfo"`
}

func main() {
	fmt.Println(FxFunc(2, 5))
	var u1, u2 uint = 15, 25
	fmt.Println(FxFunc(u1, u2))

	type UserInfo struct {
		Age int    `json:"age"`
		Sex string `json:"sex"`
	}

	// users := Users{
	// 	Name: "洛琪希",
	// 	Msg:  "🤤",
	// 	UserInfo: UserInfo{
	// 		Age: 18,
	// 		Sex: "female",
	// 	},
	// }
	// byteData, _ := json.Marshal(users)
	// fmt.Println(string(byteData))
	// 将定义的users结构体进行序列化：{"name":"洛琪希","msg":"🤤","userInfo":{"age":18,"sex":"female"}}

	// 通过将结构体设置为泛型格式来实现反序列化
	var users Users[UserInfo]
	json.Unmarshal([]byte(`{"name":"洛琪希","msg":"🤤","userInfo":{"age":18,"sex":"female"}}`), &users)
	fmt.Println(users)

	// 3.泛型切片
	type FxSlice[T int | string] []T
	Fs := FxSlice[int]{1, 3, 5, 8}
	fmt.Println(Fs)



	// 4.泛型map
	// 泛型map的key不能是any类型
	type FxMap[T int | string, K int | string] map[T]K
	Fm := FxMap[int, string]{
		1: "LBW",
	}
	fmt.Println(Fm)
}
