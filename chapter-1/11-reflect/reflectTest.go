package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

//具体类型
type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a book.")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a book.")
}

func main() {
	b := &Book{}

	var r Reader
	r = b//r实现了Reader的方法，因此被断言为Reader类型

	r.ReadBook()

	var w Writer
	w = r.(Writer)//由于r实际上是Book类型，而Book实现了Writer的方法，因此可以被断言为Writer
	w.WriteBook()
}
