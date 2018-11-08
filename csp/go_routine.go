package csp

import (
	"fmt"
	"sync"
)

/**
不用通过共享内存来通信，而是通过通信来共享内存
 */

func MainGoroutine() {
	//for i := 0; i < 10; i ++ {
	//	go Add(1, 2) //主程序不会等其他协程完成 goroutine
	//}
	//fmt.Println("add over...")

	//使用共享内存的方式
	//lock := &sync.Mutex{}
	//
	//for i := 0; i < 10; i ++ {
	//	go Count(lock)
	//}
	//for {
	//	lock.Lock()
	//	c := counter
	//	lock.Unlock()
	//	runtime.Gosched()
	//	if c >= 10 {
	//		break
	//	}
	//}

	threds := 100
	chs := make([]chan int, threds)
	for i := 0; i < threds; i++ {
		chs[i] = make(chan int)
		go CountChan(chs[i])
	}
	for _, ch := range (chs) {
		<-ch
	}

}

func Add(x, y int) {
	z := x + y
	fmt.Printf("%d add %d = %d \n", x, y, z)
}

//并发通信.有两种最常见的并发通信模型:共享数据和消息
//消息机制认为每个并发单元是自包含的、独立的个体，并且都有自己的变量，但在不同并发单元间这些变量不共享。每个并发单元的输入和输出只有一种，那就是消息

var counter int = 0

//使用共享内存的方式
func Count(lock *sync.Mutex) {
	lock.Lock()
	counter ++
	fmt.Println("counter = ", counter)
	lock.Unlock()
}

//使用channel方式来通信
func CountChan(ch chan int) {
	counter ++
	fmt.Println("Counting...counter = ", counter)
	ch <- 1

}

//基本语法 var chanName chan ElementType
//ch := make(chan int)
//写入 ch <- value  读取 value := <-ch  会阻塞程序

