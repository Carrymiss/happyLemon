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

	// 匿名函数
	test20()
	test21()
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

func test20() {
	// 匿名函数
	// 在定义匿名函数时就直接调用，这种方式匿名函数只能调用一次
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res=", res)
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test21() {
	// 匿名函数
	// 将匿名函数赋值给一个变量（函数变量），则可以通过该变量来调用该匿名函数
	a := func(n1 int, n2 int) int { // 将匿名函数赋值给变量a 类型是myFunType 自定义类型 func(int, int) int
		return n1 + n2
	}
	res := a(10, 20)
	fmt.Println("res=", res)
	// 分割线
	fmt.Println("-------------分割线--------------")
}
