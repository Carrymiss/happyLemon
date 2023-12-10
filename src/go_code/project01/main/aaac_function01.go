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
