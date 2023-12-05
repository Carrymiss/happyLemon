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
}
