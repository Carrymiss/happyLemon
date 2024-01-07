package main

import (
	"fmt"
	"time"
)

func main() {
	test110()
}

func test110() {
	// 创建2个管道
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go test111(intChan)
	go test112(intChan, exitChan)

	// 为了防止主线程退出，导致协程也退出
	// time.Sleep(time.Second * 10)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}

func test111(intChan chan int) {
	// 写入管道
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("写入数据=", i)
		time.Sleep(time.Second)
	}
	close(intChan)
}

func test112(intChan chan int, exitChan chan bool) {
	// 从管道中读取数据
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		println("读到一个数据=", v)
		time.Sleep(time.Second)
	}
	// 读取完毕后，即任务完成
	exitChan <- true
	close(exitChan)
}
