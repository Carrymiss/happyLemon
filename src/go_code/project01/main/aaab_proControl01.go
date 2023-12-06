package main

import "fmt"

func main() {
	test03()
}
func test03() {
	// switch分支
	var num int
	fmt.Println("请输入一个数字")
	fmt.Scanln(&num)
	switch num {
	case 1:
		fmt.Println("你输入的数字是1")
	case 2:
		fmt.Println("你输入的数字是2")
	case 3:
		fmt.Println("你输入的数字是3")
	default:
		fmt.Println("你输入的数字是其他数字")
	}
}
