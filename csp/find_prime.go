package csp

import (
	"fmt"
	"time"
)

/**
不用通过共享内存来通信，而是通过通信来共享内存
*/

//要找出 N 以内所有的素数,单线程进行

func MainCspPrime() {
	fmt.Println("hello SCP...")
	max := 100000
	fmt.Println("初始化切片长度", max)
	var arr []int
	for i := 2; i <= max; i++ {
		arr = append(arr, i)
	}
	//fmt.Println("arr ：", arr)
	start := time.Now().UnixNano() / 1e6
	count := selectFind(arr)
	//count := easyFind(arr)
	end := time.Now().UnixNano() / 1e6
	fmt.Println("耗时 ：", (end - start), "素数个数：", count)
}

func easyFind(arr []int) int {
	var res []int
	for _, v := range arr {
		p := true
		for j := 2; j < v; j++ {
			if v%j == 0 {
				p = false
				break
			}
		}
		if p {
			res = append(res, v)
		}
	}
	//fmt.Println("easyFind 素数共计：", count, res)
	return len(res)
}

//筛选法去除2的倍数，在去除3的倍数，以此类推
func selectFind(arr []int) int {
	var res []int

	i := 0
	for {
		//fmt.Println("循环次数 ：", i+1)
		res = remove(arr, arr[i])
		//fmt.Println("arr : ", arr, "长度 ：", len(arr))
		//fmt.Println("res : ", res, "长度 ：", len(res))
		if len(arr) == len(res) || len(res) == 1 {
			break
		}
		arr = res
		i++
	}
	//fmt.Println("selectFind返回素数：res : ", res, "共计 ：", len(res))
	return len(res)
}

//排除数组中第一个元素的倍数
func remove(arr []int, x int) []int {
	//fmt.Println("ready remove : ", arr, "x : ", x)
	var res []int
	for i := 0; i < len(arr); i++ {
		v := arr[i]
		if v == x || v%x != 0 {
			res = append(res, v)
		}
	}
	return res
}
