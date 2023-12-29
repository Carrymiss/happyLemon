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

// 工厂模式的构造函数
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说    明",
	}
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

// 将登记收入写成一个方法
func (a *FamilyAccount) income() {
	fmt.Println("本次收入金额:")
	_, _ = fmt.Scanln(&a.money)
	a.balance += a.money // 修改账户余额
	fmt.Println("本次收入说明:")
	_, _ = fmt.Scanln(&a.note)
	// 将本次收入情况,拼接到details变量
	// 收入    11000    1000    有人发红包
	//details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
	a.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", a.balance, a.money, a.note)
	a.flag = true
}

// 将登记支出写成一个方法
func (a *FamilyAccount) pay() {
	fmt.Println("本次支出金额:")
	_, _ = fmt.Scanln(&a.money)
	// 这里需要做一个必要的判断
	if a.money > a.balance {
		fmt.Println("余额的金额不足")
	}
	a.balance -= a.money
	fmt.Println("本次支出说明:")
	_, _ = fmt.Scanln(&a.note)
	a.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", a.balance, a.money, a.note)
	a.flag = true
}

// 将退出软件写成一个方法
func (a *FamilyAccount) exit() {
	fmt.Println("你确定要退出吗? y/n")
	choice := ""
	for {
		_, _ = fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误,请重新输入 y/n")
	}
	if choice == "y" {
		a.loop = false
	}
}

// 显示主菜单
func (a *FamilyAccount) mainMenu() {
	for {
		fmt.Println("----------------家庭收支记账软件----------------")
		fmt.Println("                  1 收支明细")
		fmt.Println("                  2 登记收入")
		fmt.Println("                  3 登记支出")
		fmt.Println("                  4 退出软件")
		fmt.Print("请选择(1-4):")
		_, _ = fmt.Scanln(&a.key)
		switch a.key {
		case "1":
			a.showDetails()
		case "2":
			a.income()
		case "3":
			a.pay()
		case "4":
			a.exit()
		default:
			fmt.Println("请输入正确的选项...")
		}
		if !a.loop {
			break
		}
	}
	fmt.Println("你退出了家庭记账软件的使用...")
}
