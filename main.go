package main

import (
	"fmt"
	"go-basic/oop"
)

func init() {
	fmt.Println("I am init...")
}
func main() {
	fmt.Println("Hello world,I am Golang.")

	//oop.MainInteface()
	//csp.MainGoroutine()
	//csp.MainCspPrime()
	//csp.MainCounter()
	oop.MainSlice()
}
