package main

import "fmt"

func main() {
	// 指针
	// 获取变量的地址
	var i int = 10
	fmt.Println("i的地址是", &i)

	// 创建一个指针变量
	var ptr *int = &i
	fmt.Printf("ptr的类型是%T，ptr的值是%v，ptr的地址是%v \n", ptr, ptr, &ptr)

	// 通过指针改变原数据
	var num1 int = 9
	var ptr1 *int = &num1
	*ptr1 = 10 // *ptr1代表取出ptr1指针指向的值
	fmt.Println("num1=", num1)
}
