package oop

import (
	"fmt"
)

func MainType() {
	fmt.Println("Hello OOP ,Object Oriented Programming...")
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}

	a.Add(2)
	fmt.Println("a Add 2 = ", a)

	ModifyArrayValue()
	ModifyArrayRefer()
	ModifySliceValue()
	SliceTest()
}

/*
总结：
Java语言 中值类型和对象类型
值类型：主要是基本类型，基于值语义
对象类型：可以定义成员变量和成员方法，只能通过New 在堆上创建
Any类型：整个对象类型的根，java.lang.Object. 值类型如果被Any类型引用需要装箱（Boxing）

Go语言中大多数类型的都是值语义，并且都可以包含对应的操作方法
可以给任何类型"新增"新方法。
在实现接口时候无需implement（事实上Go语言根本就不支持面向对象中的集成语法）只需实现该接口的所有方法即可。
任何类型都可以被Any类型引用，Any就是空接口 即是 interface{}
*/

//未类型添加方法，面向对象.没有Java的this 或者 python的self
//在Go语言中没有隐藏的this指针.因为方法被添加的对象显示传递，方法施加的目标不需要非得是指针
type Integer int

func (x Integer) Less(y Integer) bool {
	return x < y
}

//面向过程
func Integer_Less(x Integer, y Integer) bool {
	return x < y
}

//需要修改对象的时候传入指针，因为GoLang 类型都是基于值传递的
func (x *Integer) Add(y Integer) {
	*x += y
}

func (x Integer) PrintValue(){
	fmt.Println(x)
}

//例如 Http.Header就是一个基于map的类型，用于map的所有功能，又增加了一系列方法。 Header map[string][]string

//值语义和引用语义，GoLang中大多数都属于值语义，包括基本类型和复合类型（数组、结构体等）。
// b=a  b.modify() 如果b的修改不影响a 那么此类型属于值类型

//数组属于值类型, 表明b=a是数组内容的完整复制。
func ModifyArrayValue() {
	var a = [3]int{1, 2, 3}
	var b = a
	b[0] ++
	fmt.Println("ModifyArrayValue", a, b)
}

//想要表达引用需要使用指针,这表明b=&a赋值语句是数组内容的引用。变量b的类型不是[3]int，而是*[3]int类型。
func ModifyArrayRefer() {
	var a = [3]int{1, 2, 3}
	var b = &a
	b[0] ++
	fmt.Println("ModifyArrayRefer", a, b)
}

//切片是引用类型，因为切片内部是指向了数组指针。
func ModifySliceValue() {
	var a = []int{1, 2, 3}
	var b = a
	b[0] ++
	fmt.Println("ModifySliceValue", a, b)
}

//slice test
func SliceTest() {
	s := []int{1}
	s = append(s, 2)
	s = append(s, 3)
	x := append(s, 4)
	y := append(s, 5)
	fmt.Println("SliceTest: ", s, x, y)
}
