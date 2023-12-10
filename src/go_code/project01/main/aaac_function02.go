package main

import "fmt"

// 定义全局变量
var age = test19()

func main() {
	// init函数
	test18()
}

// init函数
func init() {
	fmt.Println("init函数执行了")
}

func test19() int {
	// 定义全局变量
	fmt.Println("全局变量执行了")
	return 90
}

func test18() {
	fmt.Println("main函数执行了")
	// 分割线
	fmt.Println("-------------分割线--------------")
}
