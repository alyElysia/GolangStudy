package main

import (
	"fmt"
	"testing"
)

func setUp() {
	fmt.Println("测试前的初始化...")
}

func tearDown() {
	fmt.Println("测试后的处理...")
}

func TestAdd2(t *testing.T) {
	t.Logf("测试失败啦~")
}

// 作为测试的入口，可在此进行测试前后的一些操作，必须以TestMain命名
func TestMain(m *testing.M) {
	setUp()
	//调用测试方法m.Run()
	m.Run()
	tearDown()
	//os.Exit(code)
}
