package test

import (
	"go-basic/ut"
	"testing"
)

func TestDivision1(t *testing.T) {
	if i, e := ut.Division(6, 2); i != 3 || e != nil {
		t.Error("除法测试没有通过")
	} else {
		t.Log("第一个测试通过了")
	}
}
