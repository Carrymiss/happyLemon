package main

import (
	"fmt"
	"happyLemon/src/go_code/project01/demo01"
	"happyLemon/src/go_code/project01/demo01/domain"
)

func main() {
	// 定义结构体
	test83()

	// 工厂模式
	test84()

	// 封装
	test85()

	// 练习
	test86()
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

func test85() {
	per := domain.NewPerson("tom")
	per.SetAge(18)
	per.SetSal(3000)
	fmt.Println(*per)
	fmt.Println("name =", per.Name, "age =", per.GetAge(), "sal =", per.GetSal())
	// 分割线
	println("-------------分割线--------------")
}

func test86() {
	account := domain.NewAccount("111113232", "123456", 3000)
	fmt.Println(*account)
	fmt.Println("accountNo =", account.AccountNo, "password =", account.Password, "balance =", account.Balance)
	// 分割线
	println("-------------分割线--------------")
}
