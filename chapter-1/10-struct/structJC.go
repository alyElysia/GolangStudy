package main

import "fmt"

// 2.面向对象特性的继承性

// 定义一个父类
type Chicken struct {
	Sex   string
	Hobby string
}

func (this *Chicken) GetSex() {
	fmt.Printf("I am %s Chicken\n", this.Sex)
}

func (this *Chicken) GetHobby() {
	fmt.Printf("Chicken love %s\n", this.Hobby)
}

// 定义一个子类
type KunKun struct {
	Chicken //表示继承该父类
	Haircut string
}

// 重写父类的方法
func (this *KunKun) GetSex() {
	fmt.Printf("I am %s \n", this.Sex)
}

func (this *KunKun) Show() {
	fmt.Println(this.Sex)
	fmt.Println(this.Hobby)
	fmt.Println(this.Haircut)
}

func main() {

	// 子类的两种定义格式
	// 1.
	k1 := KunKun{Chicken{"female", "bask"}, "zf"}
	// 2.
	var k2 KunKun
	k2.Sex = "male"
	k2.Hobby = "Jump"
	k2.Haircut = "zf"

	// 子类调用父类未被重写的方法
	k1.GetHobby()
	// 调用被子类重写的方法
	k1.GetSex()

	k1.Show()
	k2.Show()

}
