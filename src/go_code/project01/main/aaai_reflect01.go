package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 反射
	test118()
	// 结构体的反射
	test121()
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

	// 上面的rValue其实是reflect.Value类型
	// 我们可以通过rValue来获取对应的int，float32，string等等的原始类型
	num := 1 + rValue.Int()
	fmt.Println("num=", num)

	// 下面我们将rValue转成interface{}
	iV := rValue.Interface()
	// 将interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
	// 分割线
	println("-------------分割线--------------")
}

type Student struct {
	Name string
	Age  int
}

// 对结构体的反射
func test120(a interface{}) {
	rType := reflect.TypeOf(a)
	fmt.Println("rType=", rType)

	rValue := reflect.ValueOf(a)
	fmt.Println("rValue=", rValue)

	// 获取变量对应的kind
	kind1 := rType.Kind()
	kind2 := rValue.Kind()
	fmt.Printf("kind1=%v kind2=%v\n", kind1, kind2)

	iV := rValue.Interface()
	fmt.Printf("iV=%v iV的类型是%T\n", iV, iV)
	stu, ok := iV.(Student)
	if ok {
		fmt.Println("stu.Name=", stu.Name)
		fmt.Println("stu.Age=", stu.Age)
	} else {
		fmt.Println("转换失败")
	}
}

func test121() {
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	test120(stu)
}
