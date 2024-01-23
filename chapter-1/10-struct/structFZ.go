package main

import "fmt"

// 1.面向对象特性的封装性

// 若类名、属性名和方法名的首字母大写，则可以在其他包中使用
type Person struct {
	Name string
	Age  int
	Sex  string
}

// 注意！this不是关键字，它可以更换为其他可用的名称

// (this Person)的作用是将该方法与Person进行绑定
func (this Person) Show() {
	fmt.Println(this.Name)
	fmt.Println(this.Age)
	fmt.Println(this.Sex)
}

func (t *Person) GetName() string {
	return t.Name
}

func (this Person) SetAge1(age int) {
	// this指的是Person类的一个副本，无法对调用该方法的person类对象的数据进行修改
	this.Age = age
}

// 加上*之后表示的是调用的person类对象的地址，此时可以对对象进行修改
func (this *Person) SetAge(age int) {
	this.Age = age
}

func main() {
	var per Person

	per.Name = "lbw"
	per.SetAge(18)
	per.Sex = "m"
	per.Show()
}
