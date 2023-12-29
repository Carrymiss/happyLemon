package main

import "fmt"

func main() {
	displayPage()
}

func displayPage() {
	// 声明一个变量,保存接收用户输入的选项
	key := ""
	// 声明一个变量,控制是否退出for
	loop := true
	// 显示主菜单
	for {
		fmt.Println("----------------家庭收支记账软件----------------")
		fmt.Println("                  1 收支明细")
		fmt.Println("                  2 登记收入")
		fmt.Println("                  3 登记支出")
		fmt.Println("                  4 退出软件")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("收支明细")
		case "2":
			fmt.Println("登记收入")
		case "3":
			fmt.Println("登记支出")
		case "4":
			fmt.Println("退出软件")
			loop = false
		default:
			fmt.Println("请输入正确的选项")
		}
		if !loop {
			break
		}
	}
	fmt.Println("你退出家庭收支记账软件的使用...")
}
