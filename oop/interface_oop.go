package oop

import "fmt"

func MainInteface() {
	fmt.Println("Hello MainInteface...")
	var a Integer = 1
	//对象实例赋值给接口

	var b LessAdder = &a
	//var c oop.LessAdder = a //编译错误因为 Add方法属于指针方法
	var c Lesser = a
	b.Add(2)

	a.PrintValue()
	b.PrintValue()
	fmt.Println(a, b, c)

	//类型判断
	if t, ok := c.(LessAdder); ok {
		fmt.Println("var b type is", t)
	} else {
		fmt.Println("var b type is not", t)
	}

	//Any类型
	AnyType()
}

//非入侵式接口，一个累只需要实现了接口要求的所有函数我们就说这个类实现了该接口
//
type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
	PrintValue()
}

type Lesser interface {
	Less(b Integer) bool
}

//只要两个接口拥有相同的方法列表(次序不同不要紧)，那么它们就是等同的，可以相互赋值。
//如果接口A的方法列表是接口B的方法列表的子集， 那么接口B可以赋值给接口A。

//any 类型
func AnyType() {
	var v1 interface{} = 1
	var v2 interface{} = "a"
	//var v3 interface{} = &v2
	println(v1, "AnyType v1 type is")
	println("AnyType v2 type is", v2)
	//println("v1 type is", v3.(string))
}
