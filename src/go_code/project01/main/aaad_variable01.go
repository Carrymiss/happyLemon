package main

import (
	"fmt"
	"happyLemon/src/go_code/project01/demo01"
)

func main() {
	// 定义结构体
	test83()

	// 工厂模式
	test84()
}

type stundent1 struct {
	Name   string
	Age    int
	Adress string
}

func test83() {
	// 创建结构体
	// 之前定义的时候需要遵循顺序，这样就不用遵循顺序了
	var stu1 = stundent1{Name: "tom", Age: 18, Adress: "北京"}
	fmt.Println(stu1)
	// 分割线
	println("-------------分割线--------------")
}

func test84() {
	person := demo01.NewPerson("tom", 18)
	fmt.Println(*person)
	fmt.Println("name =", person.Name, "age =", person.GetAge())
	// 分割线
	println("-------------分割线--------------")
}
