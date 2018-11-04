package oop

import (
	"fmt"
)

func MainStruct() {
	fmt.Println("hello struct_oop ...")
	r1 := new(Rect)
	r2 := &Rect{}
	var r3 = &Rect{0, 0, 1, 2}
	var r4 = &Rect{width: 3, height: 4}

	fmt.Println(r1, r2, r3, r4)

	//extends 继承
	p := &Person{"张三", 18}
	s := &Student{*p, 2}
	s.SayName()
	s.SayAge()
	s.SayId()
	s.ChangeName()//注意该方法的是否为引用
	fmt.Println("after change name is ",s.Name) //如果子类中有成员变量与父类冲突，则子类覆盖父类。
}

//结构体与Java class 有同等地位，但Go 放弃了继承等特性
//Go 中没有构造函数的概念
type Rect struct {
	x, y          float64
	width, height float64
}

//可见性，首字母大写包外可见，小写包内可见，没有结构体内的范围
func (r *Rect) Area() float64 {
	return r.height * r.width
}

//父类 继承采用组合的写法
type Person struct {
	Name string
	Age  int64
}

func (b Person) SayName() {
	fmt.Println("I am ", b.Name)
}

func (b Person) SayAge() {
	fmt.Println("I am ", b.Age)
}
//指针的作用在于方法调用出是否改变
func (b *Person) ChangeName() {
	b.Name = "李四"
	fmt.Println("I change name ", b.Name)
}

//子类，如果子类中有成员变量与父类冲突，则子类覆盖父类。
type Student struct {
	Person
	Id int64
}

func (s Student) SayName() {
	fmt.Println("I am student  ", s.Name)
}

func (s Student) SayId() {
	fmt.Println("My Id is  ", s.Id)
}
