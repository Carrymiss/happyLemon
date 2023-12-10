package main

import (
	"fmt"
	"happyLemon/src/go_code/project01/demo01"
)

// 定义全局变量
var age = test19()

func main() {
	// init函数
	test18()
}

// init函数
func init() {
	demo01.Age = 100
	demo01.Name = "jack"
	fmt.Println("init函数执行了")
}

func test19() int {
	// 定义全局变量
	fmt.Println("全局变量执行了")
	return 90
}

func test18() {
	fmt.Println("main函数执行了")
	fmt.Println("age=", age)
	fmt.Println("demo01.Age=", demo01.Age)
	fmt.Println("demo01.Name=", demo01.Name)
	// 分割线
	fmt.Println("-------------分割线--------------")
}
