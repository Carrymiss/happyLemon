package main

import "fmt"

func main() {
	// test03()
	// 循环控制
	test04()
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
func test04() {
	// for循环
	for i := 1; i <= 10; i++ {
		// 在输出里面加上i，可以看到i的变化
		fmt.Println("hello world", i)
	}
	// for循环的第二种写法
	var i int = 1
	for i <= 10 {
		fmt.Println("hello world", i)
		i++
	}
}
