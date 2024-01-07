package main

import "fmt"

func main() {
	// 管道可以声明为只读或者只写
	// 1.在默认情况下，管道是双向的
	// var chan1 chan int  可读可写
	// 2.声明为只写
	test116()

	// select的使用
	//test117()
}

func test116() {
	// 声明为只写
	var chan1 chan<- int
	chan1 = make(chan int, 3)
	chan1 <- 20
	// num := <-chan1 // error
	// println("num=", num)

	// 分割线
	println("-------------分割线--------------")

	// 声明为只读
	var chan2 <-chan int
	num2 := <-chan2
	println("num2=", num2)
	// 分割线
	println("-------------分割线--------------")
}

func test117() {
	// 使用select可以解决从管道取数据的阻塞问题
	// 1.定义一个管道 10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	// 2.定义一个管道 5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}
	// 传统的方法在遍历管道时，如果不关闭会阻塞而导致deadlock
	// 问题，在实际开发中，可能我们不好确定什么时候关闭该管道
	// 可以使用select方式可以解决
	for {
		select {
		case v := <-intChan:
			println("从intChan读取的数据=", v)
		case v := <-stringChan:
			println("从stringChan读取的数据=", v)
		default:
			println("都取不到数据了")
			return
		}
	}
}
