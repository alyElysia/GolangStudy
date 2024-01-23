package main

import "fmt"

// 定义一个结构体，也就是类
type Book struct {
	author string
	name   string
}

func chBook1(book Book) {
	// 属于值传递，无法修改被调用的类对象的值
	book.author = "kjq"
}

func chBook2(book *Book) {
	// 引用传递，可以修改被调用的类对象的值
	book.author = "kjq"
}
func main() {
	var book1 Book

	fmt.Printf("the type of Book is %T\n", book1)

	book1.author = "lbw"
	book1.name = "tk"
	fmt.Println(book1)

	chBook1(book1)
	fmt.Println(book1)

	chBook2(&book1) //传递book1的地址
	fmt.Println(book1)
}
