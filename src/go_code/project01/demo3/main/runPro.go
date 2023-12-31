package main

import "happyLemon/src/go_code/project01/demo3/model"

func main() {
	// 运行项目
	customerView := model.CustomerView{
		Key:  "",
		Loop: true,
	}
	// 这里完成对customerView结构体的customerService字段的初始化
	customerView.CustomerService = model.NewCustomerService()
	// 显示主菜单
	customerView.MainMenu()
}
