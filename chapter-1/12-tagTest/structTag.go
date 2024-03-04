package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	// omitempty参数表示省略值为0的参数
	Age int `json:"age,omitempty"`
	// - 表示忽略该字段的转换
	PassWord string `json:"-"`
}

func main() {

	user := User{
		Name:     "洛琪希",
		Age:      0,
		PassWord: "15355",
	}

	userData, _ := json.Marshal(user)
	fmt.Println(string(userData))
}
