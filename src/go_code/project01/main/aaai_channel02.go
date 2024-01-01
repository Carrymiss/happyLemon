package main

import "fmt"

func main() {
	// 管道
	test107()
}

func test107() {
	// 管道
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)
	// 分割线
	println("-------------分割线--------------")
}
