package main

import (
	"fmt"
	"reflect"
)

// reflect用于获取未知变量的类型与值

// 2.获取复杂类型
type User struct {
	Id   int
	Name string
}

func (this User) Call() {
	fmt.Println("User is called ...")
}

func DoFileAndMethod(inp interface{}) {
	// 获取inp的type
	inpType := reflect.TypeOf(inp)
	fmt.Println("inpType is ", inpType)
	// 获取inp的value
	inpValue := reflect.ValueOf(inp)
	fmt.Println("inpValue is ", inpValue)

	// 通过type获取其中的字段,分为三步
	// 1.获取interface的reflect.Type,通过Type得到NumField，再进行遍历
	for i := 0; i < inpType.NumField(); i++ {
		// 2.得到每个field，它们分别表示一个数据类型
		field := inpType.Field(i)

		// 3.通过field的一个Interface()方法得到对应的value
		value := inpValue.Field(i).Interface()

		fmt.Printf("%s，%v = %v\n", field.Name, field.Type, value)
	}

	// 通过type获取其中的方法
	for i := 0; i < inpType.NumMethod(); i++ {
		m := inpType.Method(i)
		fmt.Printf("%s,%v\n", m.Name, m.Type)
	}

}

func main() {
	// 1.获取简单类型
	num := 123
	fmt.Println("num's type = ", reflect.TypeOf(num))
	fmt.Println("num's value = ", reflect.ValueOf(num))

	user := User{1,"lbw"}
	DoFileAndMethod(user)

}
