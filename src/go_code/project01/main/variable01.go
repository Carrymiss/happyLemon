package main

import "fmt"

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

}
