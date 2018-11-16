package ut

import "testing"

func TestDivision1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("除法测试没有通过")
	} else {
		t.Log("第一个测试通过了")
	}
}

func TestDivision2(t *testing.T) {
	if _, e := Division(5, 0); e == nil {
		t.Error("除数是 0, 没有抛出指定异常")
	} else {
		t.Log("第二个测试通过了")
	}
}

func TestDivision3(t *testing.T) {
	if i, e := Division(10, 5); i != 2 || e != nil {
		t.Error("除法测试 10没有通过")
	} else {
		t.Log("第三个测试通过了")
	}
}
