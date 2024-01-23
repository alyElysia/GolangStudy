package main

import (
	"encoding/json"
	"fmt"
)

// 将结构体对象编码为json对象
type Movie struct {
	Name   string   `json:"name"` // 标签的含义是将对应字段在json中的字段名
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 10, []string{"xy", "zbz"}}

	// 编码的过程	结构体 -----> json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error ....", err)
		return
	}

	fmt.Printf("jsonStr = %s\n", jsonStr)

	// 解码的过程		json  ------>   结构体
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error ... ", err)
		return
	}

	fmt.Println(myMovie)
}
