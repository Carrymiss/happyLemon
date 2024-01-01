package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {
	// 不使用协程
	test104()
	for i := 0; i < 10; i++ {
		fmt.Println("main hello,world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

	// 分割线
	println("-------------分割线--------------")

	// 设置运行的cpu数量
	test105()
}

func test104() {
	for i := 0; i < 10; i++ {
		fmt.Println("test103() hello,world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func test105() {
	// 设置运行的cpu数量
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)
	// 设置使用的cpu数量
	runtime.GOMAXPROCS(1)
	// 分割线
	println("-------------分割线--------------")
}
