package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 定义变量
	var i int
	i = 10
	fmt.Println("i=", i)

	// 声名变量的三种方式
	// 1.指定变量类型，不赋值使用默认值
	var int01 int
	fmt.Println("int01=", int01)

	// 2.类型推导
	var int02 = 10.1
	fmt.Println("int02=", int02)

	// 3.:=左侧的变量不应该是已经声明过的，否则会导致报错
	// 等价于 var name string name = "tom"
	name := "tom"
	fmt.Println("name=", name)

	// 声明多个变量
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	// 查看一个变量占用的字节大小和数据类型
	var n4 int64 = 100
	// unsafe.Sizeof(n4)是unsafe包的一个函数，可以返回n4变量占用的字节数
	fmt.Printf("n4的类型是%T，占用的字节数是%d \n", n4, unsafe.Sizeof(n4))

	// 浮点类型 可能会有精度损失
	var price float32 = 89.12
	fmt.Println("price=", price)

	// 默认的是float64
	var num1 = 1.1
	fmt.Printf("num1的数据类型是%T \n", num1)

	var num2 = .1
	fmt.Println("num2=", num2)

	// 科学计数法
	num3 := 1.2345e2  // 1.2345 * 10的2次方
	num4 := 1.2345e-2 // 1.2345 / 10的2次方
	fmt.Println("num3=", num3, "num4=", num4)

}
