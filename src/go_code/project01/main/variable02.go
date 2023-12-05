package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 类型转换
	// int -> float
	var num1 int32 = 100
	var num2 float32 = float32(num1)
	fmt.Println("num2=", num2)

	// int -> string
	var num3 int32 = 99
	var num4 string = fmt.Sprintf("%d", num3)
	fmt.Println("num4=", num4)
	// 第一个参数是要转换的值，第二个参数是转换的进制
	num6 := strconv.FormatInt(int64(num3), 10)
	fmt.Println("num6=", num6)

	// string -> int
	var str string = "123"
	var num5 int
	fmt.Sscanf(str, "%d", &num5)
	fmt.Println("num5=", num5)

}
