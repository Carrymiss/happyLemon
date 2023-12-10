package main

import (
	"fmt"
	"happyLemon/src/go_code/project01/demo01"
)

// 定义全局变量
var age = test19()

// 全局匿名函数
var (
	fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	// init函数
	test18()

	// 匿名函数
	test20()
	test21()
	test22()

	// 闭包
	test23()
	test26()

	// defer(延时机制) 关键字
	test28()
}

// init函数
func init() {
	demo01.Age = 100
	demo01.Name = "jack"
	fmt.Println("init函数执行了")
}

func test19() int {
	// 定义全局变量
	fmt.Println("全局变量执行了")
	return 90
}

func test18() {
	fmt.Println("main函数执行了")
	fmt.Println("age=", age)
	fmt.Println("demo01.Age=", demo01.Age)
	fmt.Println("demo01.Name=", demo01.Name)
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test20() {
	// 匿名函数
	// 在定义匿名函数时就直接调用，这种方式匿名函数只能调用一次
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res=", res)
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test21() {
	// 匿名函数
	// 将匿名函数赋值给一个变量（函数变量），则可以通过该变量来调用该匿名函数
	a := func(n1 int, n2 int) int { // 将匿名函数赋值给变量a 类型是myFunType 自定义类型 func(int, int) int
		return n1 + n2
	}
	res := a(10, 20)
	fmt.Println("res=", res)
	// 分割线
	fmt.Println("--------------分割线--------------")
}

func test22() {
	res := fun1(10, 20)
	fmt.Println("res=", res)
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test23() {
	// 闭包
	// 闭包是一个函数，这个函数包含了他外部作用域的一个变量
	// 底层原理：函数和他外部变量的引用（环境）
	// 闭包的说明：
	// 1.返回的是一个函数，但是这个函数引用到函数外的变量
	// 2.这个返回的函数在他被调用时，仍然可以找到他引用的外部变量
	// 3.闭包函数常见的错误用法：返回一个局部变量的地址
	// 4.闭包函数的价值：可以保存一个状态，可以反复使用这个状态（闭包函数是一个引用类型）
	f := test24()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(4))
	fmt.Println(f(5))
	// 分割线
	fmt.Println("-------------分割线--------------")
}

// 累加器
func test24() func(int) int {
	var num int
	return func(x int) int {
		num += x
		return num
	}
}

// 累加器
func test25() func() {
	var str = "hello"
	return func() {
		str += string(36)
		fmt.Println("str=", str)
	}
}

func test26() {
	// 闭包
	f := test25()
	f()
	f()
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test27(num1 int, num2 int) int {
	// defer(延时机制) 关键字 用于函数结束之前释放资源 一般用于关闭文件或者数据库连接
	// defer的执行顺序是先进后出 defer的参数会在声明时就确定
	defer fmt.Println("num1=", num1)
	defer fmt.Println("num2=", num2)
	res := num1 + num2
	fmt.Println("res=", res)
	return res
}

func test28() {
	// defer(延时机制) 关键字 用于函数结束之前释放资源 一般用于关闭文件或者数据库连接
	res := test27(10, 20)
	fmt.Println("main_res=", res)
	// 分割线
	fmt.Println("-------------分割线--------------")
}
