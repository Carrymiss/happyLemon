package model

import "fmt"

type CustomerView struct {
	// 定义必要字段
	Key  string // 接收用户输入
	Loop bool   // 表示是否循环的显示主菜单
	// 增加一个字段customerService
	CustomerService *CustomerService
}

// 显示所有的客户信息
func (cv *CustomerView) list() {
	// 首先,获取到当前所有的客户信息(在切片中)
	customers := cv.CustomerService.List()
	// 显示
	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")
}

// 添加客户的方法
func (cv *CustomerView) add() {
	fmt.Println("---------------------------添加客户---------------------------")
	fmt.Println("姓名:")
	name := ""
	_, _ = fmt.Scanln(&name)
	fmt.Println("性别:")
	gendre := ""
	_, _ = fmt.Scanln(&gendre)
	fmt.Println("年龄:")
	age := 0
	_, _ = fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	_, _ = fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email := ""
	_, _ = fmt.Scanln(&email)
	// 构建一个新的Customer实例
	customer := NewCustomer2(name, gendre, age, phone, email)
	if cv.CustomerService.Add(customer) {
		fmt.Println("---------------------------添加完成---------------------------")
	} else {
		fmt.Println("---------------------------添加失败---------------------------")
	}
}

// 删除客户的方法
func (cv *CustomerView) delete() {
	fmt.Println("---------------------------删除客户---------------------------")
	fmt.Println("请选择待删除客户编号(-1退出):")
	id := -1
	_, _ = fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除(Y/N):")
	choice := ""
	_, _ = fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if cv.CustomerService.Delete(id) {
			fmt.Println("---------------------------删除成功---------------------------")
		} else {
			fmt.Println("---------------------------删除失败---------------------------")
		}
	}
}

// MainMenu 显示主菜单
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
			cv.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			cv.list()
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
