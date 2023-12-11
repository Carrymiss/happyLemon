package main

import (
	"fmt"
	"strings"
)

func main() {
	// 如果name没有后缀，就加上后缀，否则就返回原来的名字
	test22()

	// 打印金字塔
	test23(10)

	// 打印倒金字塔
	test24(10)

	// 打印空心金字塔
	test25(10)

	// 99乘法表
	test26(9)

	// 数组转置
	test27()
}

func test21(suffix string) func(string) string {
	return func(name string) string {
		// 如果name没有后缀，就加上后缀，否则就返回原来的名字
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func test22() {
	f := test21(".jpg")
	name1 := f("test")
	name2 := f("bird.jpg")
	fmt.Println("文件名处理后：", name1)
	fmt.Println("文件名处理后：", name2)
	// 分割线
	println("-------------分割线--------------")
}

func test23(num1 int) {
	// 打印金字塔
	for i := 1; i <= num1; i++ {
		for j := 1; j <= num1-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	// 分割线
	println("-------------分割线--------------")
}

func test24(num1 int) {
	// 打印倒金字塔
	for i := 1; i <= num1; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*(num1-i)-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	// 分割线
	println("-------------分割线--------------")
}

func test25(num1 int) {
	// 打印空心金字塔
	// 最后一行隔一个打印一个*
	for i := 1; i <= num1; i++ {
		for j := 1; j <= num1-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 || (i == num1 && k%3 == 1) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	// 分割线
	println("-------------分割线--------------")
}

func test26(num int) {
	// 99乘法表
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, i*j)
		}
		fmt.Println()
	}
	// 分割线
	println("-------------分割线--------------")
}

func test27() {
	// 给一个二位数组（3*3）转置
	var arr [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var arr2 [3][3]int
	// 打印原数组
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], "\t")
		}
		fmt.Println()
	}
	// 分割线
	println("-------------分割线--------------")
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			arr2[j][i] = arr[i][j]
		}
	}
	// 打印转置后的数组
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[i]); j++ {
			fmt.Print(arr2[i][j], "\t")
		}
		fmt.Println()
	}
	// 分割线
	println("-------------分割线--------------")
}
