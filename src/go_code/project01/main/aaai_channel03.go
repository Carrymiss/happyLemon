package main

func main() {
	// 管道的关闭
	test108()

	// 遍历管道
	test109()
}

func test108() {
	// 管道的关闭
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)
	// 这时不能再写入数据到管道中
	// intChan <- 300 // panic: send on closed channel

	// 当管道关闭后，读取数据是可以的
	n1 := <-intChan
	println("n1=", n1)
	// 分割符
	println("-------------分割线--------------")
}

func test109() {
	// 遍历管道
	intChan := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan <- i * 2
	}
	// 遍历管道不能使用普通的for循环
	// for i := 0; i < len(intChan); i++ {
	// 	n1 := <-intChan
	// 	println("n1=", n1)
	// }
	// 可以使用for-range方式
	close(intChan)
	for v := range intChan {
		println("v=", v)
	}
}
