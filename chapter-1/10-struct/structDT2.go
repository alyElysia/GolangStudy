package main

import "fmt"

// interface{}是一个空接口，它被称为万能类型，当不确定传入参数的类型时可以使用该类型表示
func PrintMsg(arg interface{}) {
	fmt.Println(arg)

	// 该类型拥有断言机制，用于判断数据的类型	value表示数据的值，ok表示该数据类型是否未括号中的一致
	value, ok := arg.(string) //ok为false时value为空值，且value的类型始终为string
	if ok {
		fmt.Println(arg, "'s type is string ", value)
	} else {
		fmt.Println(arg, "'s type isn't string")
		fmt.Printf("value's type is %T\n", value)
	}

	// 可用fmt.Printf("%T\n", arg)直接获取参数的类型
	fmt.Printf("%T\n", arg)

	fmt.Println("====================")

}

type Book struct {
	T string
}

func main() {
	book := Book{"Golang"}
	num := 123
	str := "lbw"
	fla := 3.14
	PrintMsg(book)
	PrintMsg(num)
	PrintMsg(str)
	PrintMsg(fla)
}
