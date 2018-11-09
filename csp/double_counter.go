package csp

import (
	"fmt"
	"time"
)

func MainCounter() {

	//a := []int{1, 2, 3, 4, 5}
	//c := make(chan int)
	//c1 := make(chan int)
	//go sum(a[:2], c)
	//go sum(a[2:], c1)
	//x, y := <-c, <-c1 //阻塞的
	//fmt.Println("x,y = ", x, y)
	//channelbuffer()
	//c2 := make(chan int, 10)
	//go fibonacci(cap(c2), c2)
	//for i := range c2 {
	//	fmt.Println(i)
	//}
	timeoutC()

}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total //阻塞
}

func channelbuffer() {
	c := make(chan int, 5)
	c <- 1
	c <- 2
	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func fibonacci(n int, c chan int) {

	for i := 0; i < n; i++ {
		c <- i
	}
	close(c)
}
func timeoutC() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(2 * time.Second):
				println("time out")
				o <- true
				break
			}
		}
	}()

	<-o
}
