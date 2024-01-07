package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 反射
	test118()
}

func test118() {
	// 对基本数据类型 interface reflect.value进行反射的基本操作
	var num int = 100
	test119(num)
}

// 演示反射
func test119(a interface{}) {
	// 通过反射获取传入的变量的type kind 值
	// 1. 先获取到reflect.Typ
	rType := reflect.TypeOf(a)
	fmt.Println("rType=", rType)
	// 2. 获取到reflect.Value
	rValue := reflect.ValueOf(a)
	fmt.Println("rValue=", rValue)
}
