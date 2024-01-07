package main

import "fmt"

func main() {
	// 常量
	test122()
}

func test122() {
	var num int
	//常量声明的时候必须赋值 不能给默认值
	const pi = 3.1415926
	fmt.Println("num=", num, "pi=", pi)
}
