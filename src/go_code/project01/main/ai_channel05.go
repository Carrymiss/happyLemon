package main

func main() {
	// 判断是否是素数
	test113()
}

func test113() {
	intChan := make(chan int, 3)
	primeChan := make(chan int, 1000) // 放入结果
	exitChan := make(chan bool, 8)    // 退出的管道

	// 开启一个协程，向intChan放入1-8000个数
	go test114(intChan)
	// 开启8个协程，从intChan取出数据，并判断是否为素数，如果是，就放入到primeChan
	for i := 0; i < 8; i++ {
		go test115(intChan, primeChan, exitChan)
	}

	// 取出8个协程的结果
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		// 当我们从exitChan取出了8个结果，就可以放心的关闭primeChan
		close(primeChan)
	}()

	// 遍历primeChan，把结果取出
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		// 将结果输出
		println("素数=", res)
	}
	println("main线程退出")
}

// 向管道中放入1-8000个数
func test114(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

// 从intChan取出数据，并判断是否为素数，如果是，就放入到primeChan
func test115(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true // 假设是素数
		// 判断num是否为素数
		for i := 2; i < num; i++ {
			if num%i == 0 { // 说明该num不是素数
				flag = false
				break
			}
		}
		if flag {
			// 将这个数放入到primeChan
			primeChan <- num
		}
	}
	// 这里我们还不能关闭primeChan
	// 向exitChan写入true
	exitChan <- true
}
