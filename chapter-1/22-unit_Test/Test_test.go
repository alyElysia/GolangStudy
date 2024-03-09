package unit_Test

import "testing"

func TestAdd(t *testing.T) {
	res := Add(1, -1)
	if res != 1 {
		t.Errorf("出错啦~")
		return
	}

	t.Logf("测试通过拉~")
}
