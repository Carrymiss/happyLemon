package main

import (
	"fmt"
)

func main() {
	// 类型转换
	// int -> float
	var num1 int32 = 100
	var num2 float32 = float32(num1)
	fmt.Println("num2=", num2)

}
