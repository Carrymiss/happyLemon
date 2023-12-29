package main

import "fmt"

type FamilyAccount struct {
	// 声明一个变量,保存接收用户输入的选项
	key string
	// 声明一个变量,控制是否退出for
	loop bool

	// 定义账户的余额
	balance float64
	// 每次收支的金额
	money float64
	// 每次收支的说明
	note string
	// 定义一个变量,记录是否有收支的行为
	flag bool
	// 收支的详情使用字符串来记录
	// 当有收支时,只需要对details进行拼接处理即可
	details string
}

// 将显示明细写成一个方法
func (a *FamilyAccount) showDetails() {
	if a.flag {
		fmt.Println("----------------当前收支明细记录----------------")
		fmt.Println(a.details)
	} else {
		fmt.Println("当前没有收支明细... 来一笔吧!")
	}
}
