package main

import "fmt"

// 多态：同一种行为，实现方式

// 基本要素：
// 1.定义一个父类接口	本质上是一个指针
type AnimalIf interface {
	Sleep()
	GetColor() string
}

// 2.定义一个实现了父类接口所有方法的子类
type Cat struct {
	Color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleeping ...")
}

func (this *Cat) GetColor() string {
	return this.Color
}

type Dog struct {
	Color string
}

func (this *Dog) Sleep() {
	fmt.Println("Cat is sleeping ...")
}

func (this *Dog) GetColor() string {
	return this.Color
}

// 由于父类接口本质上是一个指针，因此不需再加上*
func AnimalColor(Animal AnimalIf) {
	fmt.Println("This animal's color is ", Animal.GetColor()) //这也是实现多态的一种方式
}

func main() {

	// 3.父类类型的变量（指针）指向子类类型的具体变量
	var animal AnimalIf

	// 实现多态
	animal = &Cat{"Orange"}
	animal.Sleep() //调用的是Cat实现的Sleep方法
	AnimalColor(animal)
	animal = &Dog{"Black"}
	animal.Sleep() //调用的是Cat实现的Sleep方法
	AnimalColor(animal)

}
