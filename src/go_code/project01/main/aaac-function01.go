package main

import "fmt"

func main() {
	// 函数的使用
	result := test08(1.2, 2.3, '+')
	fmt.Println("result=", result)
	result = test08(1.2, 2.3, '-')
	fmt.Println("result=", result)
	result = test08(1.2, 2.3, '*')
	fmt.Println("result=", result)
	result = test08(1.2, 2.3, '/')
	fmt.Println("result=", result)

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
