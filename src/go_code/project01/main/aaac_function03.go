package main

import "fmt"

func main() {
	// new函数
	test31()
}

func test31() {
	// new函数
	// new函数是一个内置的函数，它的函数签名如下： func new(Type) *Type
	// new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值
	// new函数分配内存，返回指针
	num := 100
	fmt.Printf("num的类型 %T num的值 %v num的地址 %v \n", num, num, &num)
	num1 := new(int)
	fmt.Printf("num1的类型 %T num1的值 %v num1的地址 %v num1指向的值 %v \n", num1, num1, &num1, *num1)
	*num1 = 100
	fmt.Printf("num1的类型 %T num1的值 %v num1的地址 %v num1指向的值 %v \n", num1, num1, &num1, *num1)
	// 分割线
	fmt.Println("-------------分割线--------------")
}
