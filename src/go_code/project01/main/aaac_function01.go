package main

import (
	"fmt"
	demo "happyLemon/src/go_code/project01/demo01" // 给导入的包起别名
)

func main() {
	// 函数的使用
	n1 := 10
	test09(n1)
	fmt.Println("main_n1=", n1)
	// 分割线
	fmt.Println("-------------分割线--------------")

	result := test08(1.2, 2.3, '+')
	fmt.Println("result=", result)
	result = test08(1.2, 2.3, '-')
	fmt.Println("result=", result)
	result = test08(1.2, 2.3, '*')
	fmt.Println("result=", result)
	result = test08(1.2, 2.3, '/')
	fmt.Println("result=", result)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 包的使用
	result = demo.Test01(1.2, 2.3, '/')
	fmt.Println("result=", result)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 递归调用
	test10(4)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 函数也是一种数据类型
	a := test11
	fmt.Printf("a的类型%T test11的类型%T\n", a, test11)
	res := a(10, 20)
	fmt.Println("res=", res)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 函数可以作为形参
	res = test12(test11, 50, 60)
	fmt.Println("res=", res)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 为了简化数据类型定义，Go支持自定义数据类型
	// 使用type关键字来定义一个自定义数据类型
	test13()
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test08(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符号错误")
	}
	return res
}

func test09(n1 int) {
	n1++
	fmt.Println("test_n1=", n1)
}

func test10(n1 int) {
	if n1 > 2 {
		n1--
		test10(n1)
	}
	fmt.Println("n1=", n1)
}

func test11(n1 int, n2 int) int {
	// 函数也是一种数据类型
	// 可以赋值给一个变量，则该变量就是一个函数类型的变量了，通过该变量可以对函数调用
	return n1 + n2
}

func test12(funVar func(int, int) int, num1 int, num2 int) int {
	// 函数可以作为形参
	return funVar(num1, num2)
}

func test13() {
	// 为了简化数据类型定义，Go支持自定义数据类型
	// 使用type关键字来定义一个自定义数据类型
	type myInt int
	var num1 myInt
	num1 = 40
	fmt.Println("num1=", num1)
	fmt.Printf("num1的类型%T\n", num1)
}
