package main

import "fmt"

func main() {
	// 求2数和
	test62()
}

type Calcuator struct {
	num1 float64
	num2 float64
}

func (cal *Calcuator) test61() float64 {
	// 求和
	return cal.num1 + cal.num2
}

func test62() {
	var cal Calcuator
	cal.num1 = 1.2
	cal.num2 = 2.3
	cal.test61()
	fmt.Printf("两数的和是：%v\n", fmt.Sprintf("%.2f", cal.test61()))
	// 分割线
	println("-------------分割线--------------")
}
