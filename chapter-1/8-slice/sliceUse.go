package main

import (
	"fmt"
)

func main() {

	// 1.切片的追加（扩容）
	slice1 := make([]int, 3, 5) //3表示切片的长度，5表示切片的容量，其中必须满足 len ≤ cap
	fmt.Printf("len = %d , cap = %d , slice1 = %d\n", len(slice1), cap(slice1), slice1)

	// 向切片添加元素
	slice1 = append(slice1, 2) //表示向切片末加上2
	fmt.Printf("len = %d , cap = %d , slice1 = %d\n", len(slice1), cap(slice1), slice1)
	slice1 = append(slice1, 2)
	// 如果对长度与容量相等的切片追加元素，会自动为切片的容量扩充最初定义的容量
	// 例如此时切片的容量已满，再追加元素后，切片的容量会加 5 ，也就是最初在make中定义的容量大小
	slice1 = append(slice1, 2)
	fmt.Printf("len = %d , cap = %d , slice1 = %d\n", len(slice1), cap(slice1), slice1)

	// tips:如果在定义切片时没有设置容量的值，则默认容量值为长度的值

	// ============================================================

	// 2.切片的截取
	slice2 := make([]int, 3)

	s2 := slice2[0:3] //截取切片的区间是左闭右开的，所以s2截取的是slice2下标为0~2区间的元素
	fmt.Printf("s2 = %d\n", s2)

	//截取区间的左侧为空表示从第一个元素开始截取，右侧区间为空表示截取到最后一个元素
	s3 := slice2[:2]
	s4 := slice2[1:]
	fmt.Printf("s3 = %d，s4 = %d\n", s3, s4)

	// 截取的切片与原切片指向的是同一个地址，因此对截取切片的元素进行修改时，原切片也会改变
	s2[0] = 110
	fmt.Printf("s2 = %d，slice2 = %d\n", s2, slice2)

	// 如果想要在不修改原切片的基础上对新切片修改，可使用copy方法
	s5 := make([]int, 3)
	copy(s5, slice2) //将slice2切片拷贝到s5

	// 对拷贝的切片修改不会影响原切片
	s5[1] = 111
	fmt.Printf("s5 = %d，slice2 = %d\n", s5, slice2)

	s6 := []int{1, 5, 2, 4}
	s6 = SliceCopy(s6, 5)
	// SliceCopy(s6, 5)
	fmt.Println(s6)
	s6 = SliceMove(s6, 4)
	// SliceMove(s6, 4)
	fmt.Println(s6)

}

// 3.切片的删除
// 虽然没有提供语法或函数，但可以根据切片的特性实现对切片中的元素进行删除

// Ⅰ.拷贝法：不会对原切片产生影响（指的是不会影响原切片的底层数组），可能会导致额外的内存分配
// a表示要删除的元素
func SliceCopy(slice []int, a int) []int {
	temp := make([]int, 0, len(slice)) //将容量设为原切片的长度，这样就不会有多余的0
	for _, v := range slice {
		if v != a {
			temp = append(temp, v)
		}
	}

	return temp
}

// Ⅱ.移位法：对原切片产生影响，不会产生额外内存分配
func SliceMove(slice []int, a int) []int {
	j := 0
	for _, v := range slice {
		if v != a {
			slice[j] = v
			j++
		}
	}

	return slice[:j]
}
