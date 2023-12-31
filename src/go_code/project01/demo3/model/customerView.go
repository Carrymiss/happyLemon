package model

import "fmt"

type CustomerView struct {
	// 定义必要字段
	Key  string // 接收用户输入
	Loop bool   // 表示是否循环的显示主菜单
	// 增加一个字段customerService
	CustomerService *CustomerService
}

// 显示主菜单
func (cv *CustomerView) MainMenu() {
	for {
		fmt.Println("----------------客户信息管理软件----------------")
		fmt.Println("                  1 添加客户")
		fmt.Println("                  2 修改客户")
		fmt.Println("                  3 删除客户")
		fmt.Println("                  4 客户列表")
		fmt.Println("                  5 退    出")
		fmt.Print("请选择(1-5):")
		_, _ = fmt.Scanln(&cv.Key)
		switch cv.Key {
		case "1":
			fmt.Println("添加客户")
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			fmt.Println("客户列表")
		case "5":
			cv.Loop = false
		default:
			fmt.Println("你的输入有误,请重新输入...")
		}
		if !cv.Loop {
			break
		}
	}
	fmt.Println("你退出了客户关系管理系统...")
}
