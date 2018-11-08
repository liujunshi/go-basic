package csp

import (
	"fmt"
	"time"
)

/**
不用通过共享内存来通信，而是通过通信来共享内存
 */

//要找出 N 以内所有的素数

func MainCsp() {
	fmt.Println("hello SCP...")
	max := 10
	fmt.Println("初始化切片长度", max)
	arr := []int{2}
	for i := 3; i <= max; i++ {
		arr = append(arr, i)
	}
	fmt.Println("arr ：", arr)
	start := time.Now().UnixNano() / 1e6
	//count := selectFind(10)
	count := easyFind(arr)
	end := time.Now().UnixNano() / 1e6
	fmt.Println("耗时 ：", (end - start), "素数个数：", count)
}

func easyFind(arr []int) int {
	count := 0
	for _, v := range arr {
		p := true
		for j := 2; j < v; j++ {
			if v%j == 0 {
				p = false
				break
			}
		}
		if p {
			count ++
			fmt.Println("is prime number : ", v, " 共计多少个：", count)
		}

	}
	return count
}

//筛选法去除2的倍数，在去除3的倍数，以此类推
func selectFind(arr []int) int {
	//for _, v := range arr {
	//	//TODO
	//}

	return len(arr)

}

func stepFind(x int, numbers *[]int) {

}
