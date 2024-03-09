package main

import "testing"

//测试项目的约定：
//	测试文件的命名：需要测试的文件名+ _test
//	测试的文件路径中不能有中文

func TestAdd(t *testing.T) {
	//1.常规测试
	res := Add(1, -1)
	if res != 0 {
		t.Errorf("出错啦~")
		return
	}
	t.Logf("测试通过拉~")

	//2.子测试
	t.Run("add1", func(t *testing.T) {
		res := Add(1, -1)
		if res != 2 {
			t.Errorf("出错啦~")
			return
		}
		t.Logf("测试通过拉~")
	})

	//3.当测试数据过多时，可将数据放在一个数组中，再遍历每组数据并通过子测试进行数据的测试
	cases := []struct {
		Name             string
		v1, v2, Excepted int
	}{
		{"add1", 1, 3, 4},
		{"add2", 2, 3, 6},
		{"add3", 1, 1, 2},
	}

	for _, entry := range cases {
		t.Run(entry.Name, func(t *testing.T) {
			r := Add(entry.v1, entry.v2)
			if r != entry.Excepted {
				t.Errorf("测试不通过~")
				return
			}

			t.Logf("测试通过啦~")
		})
	}
}
