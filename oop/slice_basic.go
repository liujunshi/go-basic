package oop

import "fmt"

func MainSlice() {
	//s := []int{0, 1, 2}
	//a := [3]int{0, 1, 2}
	//a1 := [3]int{0, 1, 2}
	//a2 := a //值类型，a2指向的是一个新的内存地址
	//a2[0] = 6
	//
	//fmt.Printf("func: %p \n", &s)
	//changeSlice(s)
	//changeArray(a)
	//changeArrayRefer(&a1)
	//
	//fmt.Println("s: ", s)
	//fmt.Println("a: ", a)
	//fmt.Println("a1: ", a1)
	//fmt.Println("a2: ", a2)

	slice()

}

//切片是引用传递？
func changeSlice(s []int) []int {
	fmt.Printf("func: %p \n", &s)
	s[1] = 111
	return s
}

//数组是类型 值传递
func changeArray(s [3]int) {
	s[0] = 5
}

//指针
func changeArrayRefer(s *[3]int) {
	s[0] = 5
}

func slice() {
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := a[:4]
	s2 := a[3:7]
	fmt.Println(s1)
	fmt.Println(s2)
	s1[3] = 100
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(a)

}
