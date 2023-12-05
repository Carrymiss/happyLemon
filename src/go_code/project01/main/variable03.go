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

	// 算数运算符
	// + - * / % ++ --
	fmt.Println(10 / 4) // 2
	var num2 float32 = 10 / 4
	fmt.Println("num2=", num2) // 2
	num3 := 10.0 / 4
	fmt.Println("num3=", num3) // 2.5

	// 取模
	fmt.Println(10 % 3)   // 1
	fmt.Println(-10 % 3)  // -1
	fmt.Println(10 % -3)  // 1
	fmt.Println(-10 % -3) // -1

	// ++ --
	var num4 int = 10
	num4++
	fmt.Println("num4=", num4) // 11
	num4--
	fmt.Println("num4=", num4) // 10

	// 关系运算符
	// == != > < >= <=
	var num5 int = 8
	var num6 int = 9
	fmt.Println(num5 == num6) // false
	fmt.Println(num5 != num6) // true
	fmt.Println(num5 > num6)  // false
	fmt.Println(num5 < num6)  // true
	fmt.Println(num5 >= num6) // false
	fmt.Println(num5 <= num6) // true

	// 逻辑运算符
	// && || !
	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}
	if age > 30 || age < 50 {
		fmt.Println("ok2")
	}
	if age > 30 && age < 50 && age != 40 {
		fmt.Println("ok3")
	}
	if !(age > 30 && age < 50) {
		fmt.Println("ok4")
	}

}
