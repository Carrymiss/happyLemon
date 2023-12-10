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
}
