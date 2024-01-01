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

	// 向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num

	// 管道的长度和容量
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan))

	// 分割线
	println("-------------分割线--------------")
}
