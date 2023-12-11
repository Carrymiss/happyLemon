package main

import (
	"errors"
	"fmt"
)

func main() {
	// new函数
	test31()

	// 异常处理
	test33()

	// 自定义错误
	test35()
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

func test32() {
	// 异常处理
	defer func() {
		err := recover() // recover()内置函数，可以捕获到异常
		if err != nil {
			fmt.Println("err=", err) // 打印异常
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func test33() {
	fmt.Println("执行命令")
	test32()
	test34()
	fmt.Println("异常之后执行命令")
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test34() {
	// 异常处理
	defer func() {
		// recover()内置函数，可以捕获到异常
		if err := recover(); err != nil {
			fmt.Println("err=", err) // 打印异常
			// 将错误发送邮件给管理员
			fmt.Println("发送邮件给管理员")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func test35() {
	// 自定义错误
	err := test36(10)
	if err != nil {
		// 如果出现错误就终止程序
		panic(err)
	}
	fmt.Println("main继续执行")
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test36(age int) (err error) {
	if age > 18 {
		return nil
	} else {
		return errors.New("年龄不合法")
	}
}
